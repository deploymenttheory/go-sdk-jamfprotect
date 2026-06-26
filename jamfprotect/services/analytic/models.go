package analytic

// Analytic represents a Jamf Protect analytic
type Analytic struct {
	UUID            string            `json:"uuid"`
	Name            string            `json:"name"`
	Label           string            `json:"label"`
	InputType       string            `json:"inputType"`
	Filter          string            `json:"filter"`
	Description     string            `json:"description"`
	LongDescription string            `json:"longDescription"`
	Created         string            `json:"created"`
	Updated         string            `json:"updated"`
	Actions         []string          `json:"actions"`
	AnalyticActions []AnalyticAction  `json:"analyticActions"`
	TenantActions   []AnalyticAction  `json:"tenantActions"`
	Tags            []string          `json:"tags"`
	Level           int               `json:"level"`
	Severity        string            `json:"severity"`
	TenantSeverity  string            `json:"tenantSeverity"`
	SnapshotFiles   []string          `json:"snapshotFiles"`
	Context         []AnalyticContext `json:"context"`
	Categories      []string          `json:"categories"`
	Jamf            bool              `json:"jamf"`
	Remediation     string            `json:"remediation"`
}

// AnalyticAction represents an action configuration for an analytic
type AnalyticAction struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// AnalyticContext represents context configuration for an analytic
type AnalyticContext struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Exprs []string `json:"exprs"`
}

// CreateAnalyticRequest is the request payload for creating an analytic
type CreateAnalyticRequest struct {
	Name            string
	InputType       string
	Description     string
	Actions         []string
	AnalyticActions []AnalyticActionInput
	Tags            []string
	Categories      []string
	Filter          string
	Context         []AnalyticContextInput
	Level           int
	Severity        string
	SnapshotFiles   []string
}

// UpdateInternalAnalyticRequest is the request payload for updating tenant-scoped fields on a managed analytic.
type UpdateInternalAnalyticRequest struct {
	TenantActions  []AnalyticActionInput
	TenantSeverity string
}

// UpdateAnalyticRequest is the request payload for updating an analytic
type UpdateAnalyticRequest struct {
	Name            string
	InputType       string
	Description     string
	Actions         []string
	AnalyticActions []AnalyticActionInput
	Tags            []string
	Categories      []string
	Filter          string
	Context         []AnalyticContextInput
	Level           int
	Severity        *string
	SnapshotFiles   []string
}

// AnalyticActionInput represents an action input for create/update
type AnalyticActionInput struct {
	Name       string
	Parameters []string
}

// AnalyticContextInput represents a context input for create/update
type AnalyticContextInput struct {
	Name  string
	Type  string
	Exprs []string
}

// ListAnalyticsResponse represents the response from listing analytics
type ListAnalyticsResponse struct {
	Items    []Analytic `json:"items"`
	PageInfo PageInfo   `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}

// AnalyticLite is a lightweight analytic summary used for listing and selection
type AnalyticLite struct {
	UUID            string   `json:"uuid"`
	Name            string   `json:"name"`
	Label           string   `json:"label"`
	InputType       string   `json:"inputType"`
	Description     string   `json:"description"`
	LongDescription string   `json:"longDescription"`
	Tags            []string `json:"tags"`
	Remediation     string   `json:"remediation"`
}

// AnalyticCategory represents an analytics category with a count
type AnalyticCategory struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// AnalyticTag represents an analytics tag with a count
type AnalyticTag struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// AnalyticsFilterOptions combines tags and categories for filter UIs
type AnalyticsFilterOptions struct {
	Tags       []AnalyticTag      `json:"tags"`
	Categories []AnalyticCategory `json:"categories"`
}

// ListAnalyticsLiteResponse is the response wrapper for listing analytics lite
type ListAnalyticsLiteResponse struct {
	Items    []AnalyticLite `json:"items"`
	PageInfo PageInfo       `json:"pageInfo"`
}
