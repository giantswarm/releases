package breakingchanges

type Finding struct {
	Severity    string `json:"severity"`    // critical, high, medium, low
	Component   string `json:"component"`   // e.g., "flatcar", "kubernetes", "cluster-azure"
	Title       string `json:"title"`       // Brief description
	Description string `json:"description"` // Detailed explanation
	Impact      string `json:"impact"`      // Who/what is affected
	Action      string `json:"action"`      // What users need to do
	Confidence  string `json:"confidence"`  // high, medium, low
	Source      string `json:"source"`      // Where this was found
	RawText     string `json:"raw_text"`    // Original text that triggered detection
}

type Release struct {
	Provider string
	Version  string
	Path     string
	README   string
	Diff     string
	YAML     string
}

type VersionChange struct {
	Component    string
	FromVersion  string
	ToVersion    string
	ChangelogURL string
	CompareURL   string
}

type ComponentMapping struct {
	RepositoryName string // The actual repository name
	ChangelogURL   string // URL template for CHANGELOG
	TagURL         string // URL template for releases
}

