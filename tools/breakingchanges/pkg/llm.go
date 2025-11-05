package breakingchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type anthropicRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []contentBlock `json:"content"`
}

type contentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// analyzeWithLLM sends the analysis context to the LLM and gets findings
func (d *Detector) analyzeWithLLM(ctx context.Context, analysisContext string) ([]Finding, error) {
	prompt := buildAnalysisPrompt(analysisContext)
	return d.callAnthropicAPI(ctx, prompt)
}

// callAnthropicAPI calls the Anthropic API with the given prompt
func (d *Detector) callAnthropicAPI(ctx context.Context, prompt string) ([]Finding, error) {
	reqBody := anthropicRequest{
		Model:     "claude-sonnet-4-5-20250929",
		MaxTokens: 4096,
		Messages: []message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", d.anthropicAPIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var apiResp anthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if len(apiResp.Content) == 0 {
		return nil, fmt.Errorf("empty response from API")
	}

	return parseFindings(apiResp.Content[0].Text)
}

// parseFindings extracts findings from the LLM response
func parseFindings(responseText string) ([]Finding, error) {
	var findings []Finding

	// Try to parse as pure JSON first
	if err := json.Unmarshal([]byte(responseText), &findings); err == nil {
		return findings, nil
	}

	// If not pure JSON, try to extract JSON from markdown code blocks
	patterns := []string{
		"```json\\s*([\\s\\S]*?)\\s*```", // ```json ... ```
		"```\\s*([\\s\\S]*?)\\s*```",     // ``` ... ```
		"\\[\\s*\\{[\\s\\S]*\\]",         // Raw JSON array
	}

	for _, pattern := range patterns {
		jsonPattern := regexp.MustCompile(pattern)
		if match := jsonPattern.FindStringSubmatch(responseText); len(match) > 0 {
			jsonStr := match[1]
			if jsonStr == "" && len(match) == 1 {
				// Third pattern - entire match is the JSON
				jsonStr = match[0]
			}
			if err := json.Unmarshal([]byte(jsonStr), &findings); err == nil {
				return findings, nil
			}
		}
	}

	// Debug: show first 500 chars of response
	debug := responseText
	if len(debug) > 500 {
		debug = debug[:500] + "..."
	}
	return nil, fmt.Errorf("failed to parse response as JSON. Response starts with: %s", debug)
}

// buildAnalysisPrompt constructs the prompt for the LLM
func buildAnalysisPrompt(context string) string {
	return fmt.Sprintf(`You are a Kubernetes and infrastructure expert analyzing a Giant Swarm workload cluster release for breaking changes.

Your task is to identify **breaking changes and potentially breaking changes** that could affect customers upgrading to this release.

**IMPORTANT**: Breaking changes are NOT limited to text containing "BREAKING" or "ACTION REQUIRED". 
You must analyze the actual code/config changes to identify things that could break customer workloads:
- **Breaking changes**: Things that WILL break or fail without customer action
- **Potentially breaking changes**: Configuration or behavior changes that MIGHT break depending on customer setup
- **High-impact changes**: Changes requiring customer awareness and likely requiring action (RBAC, security, CRDs)

## Analysis Context

%s

## Instructions

**CRITICAL - Understanding the Upgrade Path:**
The context shows version changes in "vX.Y.Z → vA.B.C" format, which means:
- The user is CURRENTLY on version vX.Y.Z (the LEFT side)
- They are UPGRADING TO version vA.B.C (the RIGHT side)
- ONLY report changes that affect moving from vX.Y.Z to vA.B.C

**Examples:**
- "Kubernetes: v1.34.1 → v1.35.0" means: upgrading FROM 1.34.1 TO 1.35.0
  - Report breaking changes NEW in Kubernetes 1.35
  - DO NOT report changes from older versions (1.32, 1.33, 1.34.0) - user already has those
- "Flatcar: v4230.2.3 → v4230.2.4" means: patch upgrade within same major version
  - Only report if there's an explicit BREAKING CHANGE note
  - DO NOT report general warnings from earlier 4230.x versions

Analyze the provided release information and identify any breaking changes. Consider:

1. **Infrastructure Changes**
   - OS-level changes (e.g., cgroups v1 removal in Flatcar)
   - Kernel changes
   - System daemon changes
   - **Check version ranges**: If a Flatcar warning mentions version 4230.2.0 but upgrade is 4230.2.3→4230.2.4, ignore it

2. **Kubernetes Changes**
   - **CRITICAL**: Look for "⚠️ URGENT UPGRADE NOTES ⚠️" section FIRST - these are mandatory action items
   - **IMPORTANT**: Read the Kubernetes changelog section and extract SPECIFIC API deprecations, removals, and behavioral changes
   - List the exact APIs being deprecated/removed (e.g., "batch/v1beta1 CronJob removed")
   - Note feature gate changes with their implications
   - Identify admission controller or RBAC changes that affect workloads
   - Pay special attention to metrics label changes, removed flags, and static pod admission changes
   - **Check context**: If upgrading 1.33.5→1.34.1, only report changes NEW in 1.34

3. **Component Changes** (Apps and Components) - Review diffs even without "BREAKING" keyword
   
   **What to look for in component diffs:**
   - **Configuration changes**: 
     - New required fields in values.yaml or ConfigMaps
     - Removed or renamed configuration options
     - Changed default values (especially for resources, replicas, timeouts)
     - New environment variables or flags
   
   - **RBAC and Security**:
     - New ClusterRole or Role permissions required
     - Changes to SecurityContext (runAsUser, capabilities, etc.)
     - New ServiceAccount requirements
     - Changes to PodSecurityPolicies or PodSecurityStandards
   
   - **CRD and API changes**:
     - New CRD versions or removed API versions
     - Required fields added to CRDs
     - Changed validation rules
     - Renamed or removed CRD fields
   
   - **Behavior changes**:
     - Changes to reconciliation logic
     - New admission webhooks or validation rules
     - Changes to default behaviors (e.g., auto-scaling, networking)
     - Migration notes or upgrade procedures mentioned
   
   - **Dependencies**:
     - Upstream chart version changes (especially major bumps)
     - New required dependencies
   
   **Version bump significance:**
   - **Major bumps** (1.x → 2.x): High chance of breaking changes, review thoroughly
   - **Minor bumps** (X.Y → X.Y+1 or X.Y+2): Often include configuration or behavior changes
   - **Patch bumps** (X.Y.Z → X.Y.Z+1): Usually safe, but still check for important fixes

4. **Dependency Changes**
   - Major version bumps that indicate breaking changes
   - Check nested dependencies (e.g., if cluster-azure uses cluster chart)
   - For Giant Swarm components (cluster-*, *-app), minor version bumps often include important changes

For each breaking or potentially breaking change, provide:

- **severity**: 
  - **critical**: WILL break customer workloads without action (removed APIs, required config fields, deleted flags)
  - **high**: LIKELY to break depending on customer setup (behavior changes, new RBAC requirements, security changes, default value changes)
  - **medium**: MIGHT break or cause issues (deprecations with migration path, CRD validation changes, new optional requirements)
  - **low**: Low risk but customers should know (minor behavior changes, informational deprecations)

- **component**: Which component is affected (e.g., "cluster-aws", "security-bundle", "Kubernetes")

- **title**: Brief one-line description (e.g., "New RBAC permissions required for cert-manager")

- **description**: Detailed explanation of what changed and why it matters

- **impact**: Who/what is affected (e.g., "All clusters using custom security policies", "Users with manual RBAC configurations")

- **action**: Specific actions users need to take, or "Review and test" if uncertain

- **confidence**: 
  - **high**: Explicitly documented or clear from code changes
  - **medium**: Likely based on version bump or configuration changes
  - **low**: Potential impact, review recommended

- **raw_text**: The exact text from the changelog/diff that triggered this finding (truncate if > 500 chars)

## DO NOT REPORT (Critical Filters)

**1. Backward Compatibility Notes for Already-Passed Versions**
   - If upgrading 1.33.5 → 1.34.1, DO NOT report notes like "affects upgrades from < 1.32" or "Before updating from Kubernetes < 1.32"
   - The user has ALREADY passed version 1.32, so these warnings don't apply

**2. Platform-Specific Changes for Unused Platforms**
   - DO NOT report Windows-related changes (Windows scheduler, Windows feature gates, HostNetworkingService)
   - This is a Linux-only environment using Flatcar Container Linux

**3. Alpha/Experimental Features Unless Explicitly Enabled**
   - DRA (Dynamic Resource Allocation) is alpha/beta - only report if there's clear evidence it's being used
   - Don't report alpha API removals unless they're widely adopted

**4. Low-Risk Patch Upgrades**
   - For small patch upgrades (e.g., cert-exporter 2.9.12 → 2.9.13), only report if there are notable changes
   - For Flatcar patch within same major (4230.2.3 → 4230.2.4), only report explicit warnings
   - Use low/medium severity for informational changes

## WHAT TO REPORT (Critical Guidance)

**FOCUS ON BREAKING CHANGES** - Report changes that could break customer workloads, even without "BREAKING" keyword:

1. **Configuration changes that could break**:
   - Required fields added (will fail if not set)
   - Fields removed or renamed (existing configs will fail)
   - Default values changed in a way that affects behavior (resources, replicas, endpoints)
   - Validation rules tightened (previously valid configs may now fail)

2. **Security and RBAC changes that could break**:
   - New permissions required (pods will fail to start without them)
   - SecurityContext changes (workloads may be denied by PSP/PSS)
   - Certificate or authentication changes (connections may fail)
   - PodSecurityStandard changes (pods may be rejected)

3. **API and CRD changes that could break**:
   - CRD API versions removed (existing resources become invalid)
   - Required fields added to CRDs (resource creation will fail)
   - Changed validation rules (existing resources may become invalid)
   - Webhooks added that could reject resources

4. **Behavior changes that could break**:
   - Changed reconciliation logic affecting existing resources
   - Changed network policies or traffic routing
   - Changed default behaviors for critical operations
   - Removed feature flags customers might be using

5. **Version jumps indicating breaking changes**:
   - Major version bumps (X.0.0 → Y.0.0) - always review
   - Components skipping minor versions (e.g., 6.2.0 → 6.4.0) - likely breaking changes
   - Minor version bumps for critical infrastructure (cluster-*, cilium, security-bundle)

**Severity assessment:**
- **Critical**: Customer workloads WILL fail without action
- **High**: Workloads LIKELY to fail based on common configurations  
- **Medium**: MIGHT break depending on specific customer setup
- **Low**: Low risk but worth documenting for customers

## Output Format

Return ONLY a JSON array of findings. Do NOT include any markdown formatting, code blocks, or explanatory text. Just pure JSON.

If no breaking changes are found, return: []

Example - BAD (too generic):
{
  "severity": "critical",
  "component": "Kubernetes",
  "title": "Kubernetes upgraded from 1.31 to 1.32",
  "description": "Minor version upgrade that typically includes API deprecations",
  "action": "Review the Kubernetes 1.32 changelog for API deprecations and removals"  ← BAD: Not specific!
}

Example - BAD (wrong version context) - upgrading 1.33.5 → 1.34.1:
{
  "severity": "critical",
  "component": "Kubernetes",
  "title": "DRA v1alpha3 API removed",
  "description": "The v1beta1 API was introduced in 1.32",
  "action": "Before updating from Kubernetes < 1.32, delete all resources"  ← BAD: User is already on 1.33.5!
}

Example - BAD (Windows on Linux cluster):
{
  "severity": "medium",
  "component": "Kubernetes",
  "title": "Windows scheduler profile removed"  ← BAD: This is a Linux-only cluster!
}

Example - BAD (patch upgrade noise) - upgrading 4230.2.3 → 4230.2.4:
{
  "severity": "medium",
  "component": "Flatcar",
  "title": "CGroups V1 deprecated in 4230.2.0"  ← BAD: This is a patch upgrade, 4230.2.0 is already passed!
}

Example - GOOD (specific Kubernetes change):
{
  "severity": "high",
  "component": "Kubernetes",
  "title": "flowcontrol.apiserver.k8s.io/v1beta3 API removed",
  "description": "The flowcontrol.apiserver.k8s.io/v1beta3 API version is removed in 1.34. Use v1 instead.",
  "impact": "Manifests using v1beta3 FlowSchema or PriorityLevelConfiguration will fail",
  "action": "Update all FlowSchema and PriorityLevelConfiguration manifests to use flowcontrol.apiserver.k8s.io/v1",
  "confidence": "high",
  "raw_text": "Removed flowcontrol.apiserver.k8s.io/v1beta3 API"
}

Example - GOOD (component configuration change - not breaking, but important):
{
  "severity": "medium",
  "component": "cluster-aws",
  "title": "Default instance type changed from t3.medium to t3.small",
  "description": "cluster-aws v6.4.0 changes the default worker node instance type",
  "impact": "New clusters will use smaller instances unless explicitly configured",
  "action": "Review and update instance type configuration if t3.medium is required",
  "confidence": "high",
  "raw_text": "Changed default instanceType from t3.medium to t3.small in values.yaml"
}

Example - GOOD (RBAC change - customers should be aware):
{
  "severity": "high",
  "component": "cert-manager",
  "title": "New ClusterRole permissions required for certificate management",
  "description": "cert-manager 3.9.4 adds new 'certificatesigningrequests/approval' permission to ClusterRole",
  "impact": "Clusters with custom RBAC configurations may need manual permission updates",
  "action": "Verify cert-manager has required permissions if using custom RBAC",
  "confidence": "high",
  "raw_text": "Added certificatesigningrequests/approval to ClusterRole rules"
}

Example - GOOD (security-bundle policy change - informational but important):
{
  "severity": "low",
  "component": "security-bundle",
  "title": "Kyverno policy updated to enforce Pod Security Standards",
  "description": "security-bundle 1.14.0 updates policies to align with Kubernetes v1.34 Pod Security Standards",
  "impact": "Improved security posture; workloads already compliant should not be affected",
  "action": "Review and test workloads if using custom security contexts",
  "confidence": "medium",
  "raw_text": "Updated kyverno policies to enforce PSS restricted profile"
}

IMPORTANT: Your response must be ONLY the JSON array, nothing else. No markdown, no code blocks, no explanations.`, context)
}
