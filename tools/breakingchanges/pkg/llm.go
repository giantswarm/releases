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

Your task is to identify potential breaking changes that could affect customers upgrading to this release.

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
   - **IMPORTANT**: Read the Kubernetes changelog section and extract SPECIFIC API deprecations, removals, and behavioral changes
   - List the exact APIs being deprecated/removed (e.g., "batch/v1beta1 CronJob removed")
   - Note feature gate changes with their implications
   - Identify admission controller or RBAC changes that affect workloads
   - **Check context**: If upgrading 1.33.5→1.34.1, only report changes NEW in 1.34

3. **Component Changes**
   - Configuration flag removals
   - Default behavior changes
   - Deprecated feature removals
   - Required permission changes

4. **Dependency Changes**
   - Major version bumps that indicate breaking changes
   - Check nested dependencies (e.g., if cluster-azure uses cluster chart)
   - **Patch upgrades** (X.Y.Z → X.Y.Z+1) rarely have breaking changes

For each breaking change found, provide:
- **severity**: critical (requires action before upgrade) | high (likely needs action) | medium (review recommended) | low (minimal impact)
- **component**: Which component is affected
- **title**: Brief one-line description
- **description**: Detailed explanation of what changed
- **impact**: Who/what is affected
- **action**: Specific actions users need to take
- **confidence**: high (explicitly stated as breaking) | medium (likely breaking) | low (potentially breaking)
- **raw_text**: The exact text from the changelog that triggered this finding

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

**4. Patch-Level Upgrades**
   - For X.Y.Z → X.Y.Z+1 upgrades (e.g., 4230.2.3 → 4230.2.4), breaking changes are extremely rare
   - Only report if there's an explicit "BREAKING" or "ACTION REQUIRED" note

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

Example - GOOD (specific from changelog):
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

IMPORTANT: Your response must be ONLY the JSON array, nothing else. No markdown, no code blocks, no explanations.`, context)
}
