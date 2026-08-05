// Publishes the `E2E Coverage` check run for release PRs in this repo.
//
// Requires the full set of E2E test suites for every newly added release, whatever
// release stage the PR is currently in, so a release can't be merged having only run
// the suites its stage happens to trigger. See giantswarm/roadmap#4334 and the
// "Test coverage requirements" section of the README.
//
// Run from .github/workflows/e2e-coverage.yaml. Uses the REST API directly rather than
// actions/github-script: pinning that action's SHA trips the repo's gitleaks rule,
// which matches any long alphanumeric run next to the word "github".
//
// Environment:
//   GITHUB_TOKEN       - token with checks:write
//   GITHUB_REPOSITORY  - owner/repo
//   GITHUB_EVENT_NAME  - triggering event
//   GITHUB_EVENT_PATH  - path to the event payload
//   CAPI_NAMES         - JSON map of provider directory → CAPI name
//   PR_NUMBER          - PR to evaluate (workflow_dispatch only)

const fs = require('fs')

const API_ROOT = 'https://api.github.com'
const [owner, repo] = (process.env.GITHUB_REPOSITORY || '').split('/')
const token = process.env.GITHUB_TOKEN
const eventName = process.env.GITHUB_EVENT_NAME
const payload = process.env.GITHUB_EVENT_PATH && fs.existsSync(process.env.GITHUB_EVENT_PATH)
  ? JSON.parse(fs.readFileSync(process.env.GITHUB_EVENT_PATH, 'utf8'))
  : {}

const log = (message) => console.log(message)

// Single request against the repo's REST namespace. `path` is relative to
// /repos/{owner}/{repo}, e.g. `/pulls/2400`.
const request = async (path, method = 'GET', body = undefined) => {
  const url = path.startsWith('http') ? path : `${API_ROOT}/repos/${owner}/${repo}${path}`
  const response = await fetch(url, {
    method,
    headers: {
      accept: 'application/vnd.github+json',
      authorization: `Bearer ${token}`,
      'content-type': 'application/json',
      'user-agent': 'giantswarm-releases-e2e-coverage',
      'x-github-api-version': '2022-11-28',
    },
    body: body === undefined ? undefined : JSON.stringify(body),
  })

  if (!response.ok) {
    throw new Error(`${method} ${url} failed: ${response.status} ${await response.text()}`)
  }
  return { data: response.status === 204 ? null : await response.json(), response }
}

const api = async (path, method = 'GET', body = undefined) => (await request(path, method, body)).data

// Follows Link headers and concatenates the pages. Used for list endpoints that
// return a bare array.
const paginate = async (path, extract = (data) => data) => {
  let next = `${API_ROOT}/repos/${owner}/${repo}${path}${path.includes('?') ? '&' : '?'}per_page=100`
  const items = []
  while (next) {
    const { data, response } = await request(next)
    items.push(...extract(data))
    const link = response.headers.get('link') || ''
    const match = link.match(/<([^>]+)>;\s*rel="next"/)
    next = match ? match[1] : null
  }
  return items
}

// Check runs come back wrapped in an object rather than as a bare array.
const listCheckRuns = (ref, checkName) => paginate(
  `/commits/${ref}/check-runs${checkName ? `?check_name=${encodeURIComponent(checkName)}` : ''}`,
  (data) => data.check_runs,
)

const main = async () => {
  const CHECK_NAME = 'E2E Coverage'
  const SUITE_CHECK_PREFIX = 'Release Tests / '

  // Directory name → CAPI name used in the per-suite check runs (from get-providers).
  const capiNames = JSON.parse(process.env.CAPI_NAMES)

  // The full ("freeze stage") matrix that releases-test-suites expands to per
  // provider, keyed by the provider directory in this repo.
  const EXPECTED_SUITES = {
    'capa': ['standard', 'upgrade', 'upgrade-major', 'cilium-eni-mode', 'china', 'private'],
    'azure': ['standard', 'upgrade', 'upgrade-major', 'private'],
    'vsphere': ['standard', 'upgrade', 'upgrade-major', 'on-capa', 'on-capz'],
    'cloud-director': ['standard', 'upgrade', 'upgrade-major'],
    'eks': ['standard', 'upgrade', 'upgrade-major'],
    // aks and proxmox have no releases-test-suites mapping yet, so nothing is expected.
    'aks': [],
    'proxmox': [],
  }

  // Suite → the check run name produced by check-run-results-to-pr, minus the
  // prefix. `{capi}` is the provider's CAPI name, e.g. CAPA.
  const CHECK_NAME_FORMATS = {
    'standard': '{capi} Standard Suite',
    'upgrade': '{capi} Upgrade Suite',
    'upgrade-major': '(Previous Major) {capi} Upgrade Suite',
    'upgrade-major-first': '(First Previous Major) {capi} Upgrade Suite',
    'private': '{capi} Private Suite',
    'china': '{capi} China Suite',
    'cilium-eni-mode': '{capi} Cilium ENI Mode Suite',
    'on-capa': '{capi} on CAPA Suite',
    'on-capz': '{capi} on CAPZ Suite',
  }

  // Provider directory → the provider name used in cluster-test-suites paths.
  const SUITE_PATH_PROVIDER = {
    'capa': 'capa', 'azure': 'capz', 'vsphere': 'capv', 'cloud-director': 'capvcd', 'eks': 'eks',
  }

  // Waivers may be written with either spelling, e.g. `azure/private` or `capz/private`.
  const PROVIDER_ALIASES = { 'aws': 'capa' }
  for (const [dir, suiteProvider] of Object.entries(SUITE_PATH_PROVIDER)) {
    PROVIDER_ALIASES[dir] = dir
    PROVIDER_ALIASES[suiteProvider] = dir
  }

  const FAILED_CONCLUSIONS = ['failure', 'timed_out', 'cancelled', 'action_required', 'stale']


  // --- Resolve the PR this run is about -----------------------------------------
  let prNumber
  switch (eventName) {
    case 'pull_request':
      prNumber = payload.pull_request.number
      break
    case 'issue_comment':
      prNumber = payload.issue.number
      break
    case 'workflow_dispatch':
      prNumber = parseInt(process.env.PR_NUMBER, 10)
      break
    case 'check_run': {
      const attached = payload.check_run.pull_requests || []
      if (attached.length > 0) {
        prNumber = attached[0].number
      } else {
        // Check runs created through the API aren't always linked to their PR.
        const associated = await api(`/commits/${payload.check_run.head_sha}/pulls`)
        const open = associated.find(candidate => candidate.state === 'open')
        if (!open) {
          log('Check run is not associated with an open PR, nothing to do')
          return
        }
        prNumber = open.number
      }
      break
    }
    default:
      log(`Unsupported event ${eventName}, nothing to do`)
      return
  }

  const pr = await api(`/pulls/${prNumber}`)
  if (pr.state !== 'open') {
    log(`PR #${prNumber} is ${pr.state}, skipping`)
    return
  }
  if (pr.head.repo.full_name !== `${owner}/${repo}`) {
    log('PR is from a fork, skipping - the token cannot write check runs')
    return
  }

  const headSha = pr.head.sha
  log(`Evaluating coverage for PR #${prNumber} at ${headSha}`)

  // The result is published twice, because the two mechanisms are good at different
  // things:
  //
  //   * A check run carries the full per-provider report, but GitHub adopts API-created
  //     check runs into an existing check suite for this app and commit, so it ends up
  //     grouped under an unrelated workflow (`gitleaks / E2E Coverage`) and is easy to
  //     miss in the Checks tab.
  //   * A commit status is always its own top-level row - like the `ci/circleci: ...`
  //     rows - so it is easy to find, and it can be required in branch protection
  //     directly rather than only through pr-gatekeeper. It carries no markdown, so its
  //     `target_url` points at the check run holding the detail.
  //
  // The check run is updated in place rather than recreated, because pr-gatekeeper
  // refuses to evaluate a name that matches several runs on one commit.
  const publish = async ({ status, conclusion, title, summary }) => {
    const existing = await listCheckRuns(headSha, CHECK_NAME)
    const checkPayload = { status, output: { title, summary } }
    if (conclusion) checkPayload.conclusion = conclusion

    let checkRun
    if (existing.length > 0) {
      const newest = existing.sort((a, b) => new Date(b.started_at) - new Date(a.started_at))[0]
      checkRun = await api(`/check-runs/${newest.id}`, 'PATCH', checkPayload)
    } else {
      checkRun = await api('/check-runs', 'POST', { name: CHECK_NAME, head_sha: headSha, ...checkPayload })
    }

    // Statuses have no in-progress state; `pending` is the equivalent and blocks just the
    // same. Descriptions are limited to 140 characters.
    const state = conclusion === 'success' ? 'success' : (conclusion ? 'failure' : 'pending')
    await api(`/statuses/${headSha}`, 'POST', {
      context: CHECK_NAME,
      state,
      description: title.length > 140 ? `${title.slice(0, 137)}...` : title,
      target_url: checkRun && checkRun.html_url ? checkRun.html_url : undefined,
    })
    log(`Published status ${CHECK_NAME}=${state}`)
  }

  // --- Which new releases does this PR add? -------------------------------------
  // Only newly added releases need coverage. Deprecating (modified) or archiving
  // (renamed) an existing release leaves its release.yaml otherwise untouched.
  const files = await paginate(`/pulls/${prNumber}/files`)

  const releaseYamlPattern = /^([^/]+)\/(v[0-9]+\.[0-9]+\.[0-9]+[^/]*)\/release\.yaml$/
  const releases = new Map() // provider dir → Set of versions
  const releaseYamlPaths = new Set()
  for (const file of files) {
    const match = file.filename.match(releaseYamlPattern)
    if (!match || file.status !== 'added') continue
    const [, provider, version] = match
    if (!(provider in EXPECTED_SUITES)) continue
    if (!releases.has(provider)) releases.set(provider, new Set())
    releases.get(provider).add(version)
    releaseYamlPaths.add(file.filename)
  }

  if (releases.size === 0) {
    log('PR adds no new releases, reporting coverage as not applicable')
    await publish({
      status: 'completed',
      conclusion: 'success',
      title: 'No new releases added',
      summary: 'This PR does not add a new release, so no E2E test coverage is required.',
    })
    return
  }

  // --- Coverage stays valid until the release content changes -------------------
  // Walk the PR's commits newest first and stop at the first one whose release
  // content differs from head. Suites that passed on the commits in between still
  // count, so README/announcement-only commits don't invalidate coverage.
  const commits = await paginate(`/pulls/${prNumber}/commits`)

  // devctl sets spec.date to the current time on every generation (see
  // devctl pkg/release/create.go), so even an /update-release that picks up no new
  // versions leaves a diff. A timestamp is not something the tests can observe, so a
  // date-only diff must not invalidate results. `spec.state` is deliberately not
  // ignored: it is a meaningful change, unlike a generated timestamp.
  const isDateOnlyDiff = (patch) => {
    // No patch means the diff was too large to include - assume the content changed.
    if (!patch) return false
    const changedLines = patch
      .split('\n')
      .filter(line => /^[+-]/.test(line) && !/^(\+\+\+|---)/.test(line))
    return changedLines.length > 0 && changedLines.every(line => /^[+-]\s*date:\s*"/.test(line))
  }

  const releaseContentChanged = async (sha) => {
    const comparison = await api(`/compare/${sha}...${headSha}`)
    return (comparison.files || [])
      .filter(file => releaseYamlPaths.has(file.filename))
      .some(file => !isDateOnlyDiff(file.patch))
  }

  const validShas = []
  let contentChangedAt = null
  for (const commit of [...commits].reverse()) {
    if (commit.sha !== headSha && await releaseContentChanged(commit.sha)) {
      // This commit, and anything older, predates the current release content.
      contentChangedAt = new Date(commit.commit.committer.date)
      break
    }
    validShas.push(commit.sha)
  }
  if (!validShas.includes(headSha)) validShas.unshift(headSha)
  log(`Counting results from ${validShas.length} commit(s) holding the current release content`)

  // --- Collect suite results across those commits -------------------------------
  // validShas is ordered newest first, and within one commit the most recently
  // started run wins. A success is then kept even if another run of the same suite
  // failed afterwards: every commit considered here carries identical release
  // content, so the pass still applies to what is being merged. A genuinely red
  // run is still surfaced by the `Release Test Suites` check, which the PR
  // gatekeeper requires separately.
  const suiteResults = new Map() // check name → { conclusion, url, sha }
  for (const sha of validShas) {
    const checkRuns = await listCheckRuns(sha)
    const byStartedAt = checkRuns
      .filter(run => run.name.startsWith(SUITE_CHECK_PREFIX))
      .sort((a, b) => new Date(b.started_at) - new Date(a.started_at))

    for (const run of byStartedAt) {
      const existing = suiteResults.get(run.name)
      if (existing && (existing.conclusion === 'success' || existing.sha === sha)) continue
      suiteResults.set(run.name, { conclusion: run.conclusion, url: run.html_url, sha })
    }
  }

  // --- Waivers ------------------------------------------------------------------
  // `/waive-suite <provider>/<suite> <reason>` from an org member waives one suite.
  // A waiver is void once the release content changes, same as a test result.
  const comments = await paginate(`/issues/${prNumber}/comments`)

  const waivers = new Map() // `${providerDir}/${suite}` → { user, reason }
  const waiverPattern = /^\/waive-suite\s+(?:\.\/providers\/)?([a-z-]+)\/([a-z0-9-]+)\s+(.+)$/i
  for (const comment of comments) {
    if (!['OWNER', 'MEMBER', 'COLLABORATOR'].includes(comment.author_association)) continue
    if (contentChangedAt && new Date(comment.created_at) < contentChangedAt) continue
    for (const line of (comment.body || '').split(/\r?\n/)) {
      const match = line.trim().match(waiverPattern)
      if (!match) continue
      const [, rawProvider, suite, reason] = match
      const provider = PROVIDER_ALIASES[rawProvider.toLowerCase()]
      if (!provider) continue
      waivers.set(`${provider}/${suite.toLowerCase()}`, {
        user: comment.user.login,
        reason: reason.trim(),
      })
    }
  }

  // --- What would a plain `/run releases-test-suites` cover right now? ----------
  // releases-test-suites picks suites from the stage label, so the trigger to
  // suggest depends on the stage: the variant suites only run at freeze.
  const stageLabel = ['freeze', 'active', 'development']
    .find(stage => pr.labels.some(label => label.name === `stage/${stage}`)) || ''

  const suitesRunByStage = (provider) => {
    const all = EXPECTED_SUITES[provider]
    // No stage label means the pipeline runs the full matrix.
    if (stageLabel === 'freeze' || stageLabel === '') return all
    const base = stageLabel === 'development'
      ? ['standard', 'upgrade', 'upgrade-major']
      : ['standard', 'upgrade']
    return all.filter(suite => base.includes(suite))
  }

  // release-stages.yaml only allows development -> active -> freeze, so the next
  // command depends on where the PR is now. Suggesting `/stage freeze` from
  // development would be rejected as an invalid transition.
  const NEXT_STAGE_COMMAND = { '': '/stage active', 'development': '/stage active', 'active': '/stage freeze' }
  const stageAdvanceCommand = NEXT_STAGE_COMMAND[stageLabel]
  const stageRunsDescription = {
    'development': '`standard`, `upgrade` and `upgrade-major` per provider',
    'active': '`standard` and `upgrade` per provider',
    'freeze': 'every suite per provider',
    '': 'every suite per provider',
  }[stageLabel]

  // --- Evaluate -----------------------------------------------------------------
  const lines = []
  const outstanding = [] // { provider, suite, path } still to pass
  let expectedCount = 0
  let coveredCount = 0
  let anyFailed = false

  const byCapiName = (a, b) => (capiNames[a[0]] || a[0]).localeCompare(capiNames[b[0]] || b[0])
  for (const [provider, versions] of [...releases.entries()].sort(byCapiName)) {
    const capi = capiNames[provider] || provider
    const suites = EXPECTED_SUITES[provider]
    const versionList = [...versions].sort().join(', ')

    if (suites.length === 0) {
      lines.push(`### ${capi} (${versionList})`, '', 'ℹ️ No E2E test suites are defined for this provider - nothing to enforce.', '')
      continue
    }

    lines.push(`### ${capi} (${versionList})`, '', '| Suite | Status |', '| --- | --- |')

    for (const suite of suites) {
      expectedCount++
      const suitePath = `./providers/${SUITE_PATH_PROVIDER[provider]}/${suite}`
      const checkName = SUITE_CHECK_PREFIX + CHECK_NAME_FORMATS[suite].replace('{capi}', capi)
      const result = suiteResults.get(checkName)
      const waiver = waivers.get(`${provider}/${suite}`)

      if (result && result.conclusion === 'success') {
        coveredCount++
        const note = result.sha === headSha
          ? ''
          : ` (on \`${result.sha.substring(0, 7)}\`, release content unchanged since)`
        lines.push(`| [${suite}](${result.url}) | ✅ passed${note} |`)
      } else if (waiver) {
        coveredCount++
        lines.push(`| ${suite} | ⚠️ waived by @${waiver.user} - _${waiver.reason}_ |`)
      } else if (result && FAILED_CONCLUSIONS.includes(result.conclusion)) {
        anyFailed = true
        outstanding.push({ provider, suite, path: suitePath })
        lines.push(`| [${suite}](${result.url}) | ❌ ${result.conclusion} |`)
      } else if (result) {
        outstanding.push({ provider, suite, path: suitePath })
        lines.push(`| ${suite} | 🚧 in progress |`)
      } else {
        outstanding.push({ provider, suite, path: suitePath })
        // Make it obvious which suites the current stage will never start on
        // its own, so a missing result doesn't look like something went wrong.
        const stageNote = suitesRunByStage(provider).includes(suite)
          ? ''
          : ' - runs automatically at `stage/freeze`'
        lines.push(`| ${suite} | ⬜ not run${stageNote} |`)
      }
    }
    lines.push('')
  }

  const complete = coveredCount === expectedCount

  if (!complete) {
    // While most of what the current stage runs is still outstanding, the plain
    // trigger is the better advice: it covers all of them and keeps the hint
    // short. Once the stage's suites are mostly green, re-running them all just
    // burns test capacity, so list only what is actually outstanding instead.
    const runByStage = (entry) => suitesRunByStage(entry.provider).includes(entry.suite)
    const stageTotal = [...releases.keys()].reduce((total, provider) => total + suitesRunByStage(provider).length, 0)
    const stageOutstanding = outstanding.filter(runByStage)
    const notRunByStage = outstanding.filter(entry => !runByStage(entry))
    const stageRerunCount = stageTotal - stageOutstanding.length
    const useStageTrigger = stageTotal > 0 && stageOutstanding.length * 2 > stageTotal

    const stageDescription = stageLabel ? `\`stage/${stageLabel}\`` : 'this PR (no stage label)'

    // Stages advance one step at a time, so spell out the command that applies now
    // and mention the step after it when freeze is still two transitions away.
    const stageAdvanceInstruction = stageAdvanceCommand
      ? (stageAdvanceCommand === '/stage active'
          ? `Advance the stage with \`/stage active\` when the release is ready for team review, then \`/stage freeze\` - stages move one step at a time.`
          : `Advance the stage with \`/stage freeze\` when the release is ready.`)
      : ''

    lines.push('---', '', '#### Run the outstanding suites', '')

    if (useStageTrigger) {
      const rerunNote = stageRerunCount > 0
        ? ` It also re-runs the ${stageRerunCount} suite(s) that already passed - target them individually with \`TARGET_SUITES\` if you want to avoid that.`
        : ''
      lines.push(
        `This runs every suite ${stageDescription} covers${notRunByStage.length > 0 ? '' : ', which is all of the outstanding ones'}.${rerunNote}`,
        '',
        '```',
        '/run releases-test-suites',
        '```',
        '',
      )
      if (notRunByStage.length > 0) {
        lines.push(
          `The remaining ${notRunByStage.length} suite(s) are not run by ${stageDescription} and only start automatically once the PR reaches \`stage/freeze\`. ${stageAdvanceInstruction} Or run them now without changing the stage:`,
          '',
          '```',
          `/run releases-test-suites TARGET_SUITES=${notRunByStage.map(entry => entry.path).join(',')}`,
          '```',
          '',
        )
      }
    } else {
      lines.push(
        '```',
        `/run releases-test-suites TARGET_SUITES=${outstanding.map(entry => entry.path).join(',')}`,
        '```',
        '',
      )
      if (notRunByStage.length > 0 && stageLabel && stageLabel !== 'freeze') {
        lines.push(
          `Note that ${notRunByStage.length} of these are not run by ${stageDescription}, so a plain \`/run releases-test-suites\` will not cover them until the PR reaches \`stage/freeze\`. ${stageAdvanceInstruction}`,
          '',
        )
      }
    }

    lines.push(
      'If a suite genuinely cannot pass - for example because its test environment is down -',
      'waive it with a reason. The waiver is recorded here and becomes void as soon as a',
      'release in this PR changes:',
      '',
      '```',
      '/waive-suite capa/china Beijing environment down, see giantswarm/roadmap#1234',
      '```',
      '',
    )
  }

  lines.push(
    '---',
    '_Coverage is required regardless of the release stage - see [giantswarm/roadmap#4334](https://github.com/giantswarm/roadmap/issues/4334)._',
  )

  // Put the stage in the title so it is readable from the checks list without
  // expanding anything - it is the main reason suites are missing.
  const stageSuffix = stageLabel ? ` · stage/${stageLabel}` : ''
  const title = complete
    ? `All ${expectedCount} expected suites covered${stageSuffix}`
    : `${coveredCount}/${expectedCount} expected suites covered${stageSuffix}`

  const stageHeader = stageLabel
    ? `**Current stage:** \`stage/${stageLabel}\` - a plain \`/run releases-test-suites\` runs ${stageRunsDescription}.`
    : `**Current stage:** none - a plain \`/run releases-test-suites\` runs ${stageRunsDescription}.`

  const summary = complete
    ? `✅ Every expected E2E test suite has passed for the releases added in this PR.\n\n${stageHeader}`
    : `🚧 ${expectedCount - coveredCount} expected E2E test suite(s) still need to pass before this release can be merged.\n\n${stageHeader}`

  // A failed suite is reported as a failure. Suites that simply haven't run yet
  // leave the check in progress, which blocks the merge without marking the PR
  // as broken while a release is still being worked on.
  await publish({
    status: complete || anyFailed ? 'completed' : 'in_progress',
    conclusion: complete ? 'success' : (anyFailed ? 'failure' : undefined),
    title,
    summary: `${summary}\n\n${lines.join('\n')}`,
  })

  log(`${title} (${complete ? 'success' : anyFailed ? 'failure' : 'in progress'})`)
}

const finished = main()
finished.catch((error) => {
  console.error(error)
  process.exitCode = 1
})

module.exports = finished
