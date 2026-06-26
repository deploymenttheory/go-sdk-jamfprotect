package plan

// Plan represents a Jamf Protect plan
type Plan struct {
	ID                   string                `json:"id"`
	Hash                 string                `json:"hash"`
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	Created              string                `json:"created"`
	Updated              string                `json:"updated"`
	LogLevel             string                `json:"logLevel"`
	AutoUpdate           bool                  `json:"autoUpdate"`
	CommsConfig          *CommsConfig          `json:"commsConfig"`
	InfoSync             *InfoSync             `json:"infoSync"`
	SignaturesFeedConfig *SignaturesFeedConfig `json:"signaturesFeedConfig"`
	ActionConfigs        *PlanRef              `json:"actionConfigs"`
	ExceptionSets        []ExceptionSet        `json:"exceptionSets"`
	USBControlSet        *PlanRef              `json:"usbControlSet"`
	Telemetry            *PlanRef              `json:"telemetry"`
	TelemetryV2          *PlanRef              `json:"telemetryV2"`
	AnalyticSets         []AnalyticSet         `json:"analyticSets"`
}

// CommsConfig represents communications configuration in a plan
type CommsConfig struct {
	FQDN     string `json:"fqdn"`
	Protocol string `json:"protocol"`
}

// InfoSync represents info sync configuration in a plan
type InfoSync struct {
	Attrs                []string `json:"attrs"`
	InsightsSyncInterval int64    `json:"insightsSyncInterval"`
}

// SignaturesFeedConfig represents signatures feed configuration in a plan
type SignaturesFeedConfig struct {
	Mode string `json:"mode"`
}

// PlanRef represents an entity reference in a plan
type PlanRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExceptionSet represents an exception set in a plan
type ExceptionSet struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Managed bool   `json:"managed"`
}

// AnalyticSet represents an analytic set in a plan
type AnalyticSet struct {
	Type        string         `json:"type"`
	AnalyticSet AnalyticSetRef `json:"analyticSet"`
}

// AnalyticSetRef represents an analytic set reference
type AnalyticSetRef struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Managed   bool       `json:"managed"`
	Analytics []Analytic `json:"analytics"`
}

// Analytic represents analytic metadata on a plan analytic set
type Analytic struct {
	UUID       string   `json:"uuid"`
	Categories []string `json:"categories"`
}

// CreatePlanRequest is the request payload for creating a plan
type CreatePlanRequest struct {
	Name                 string
	Description          string
	LogLevel             *string
	ActionConfigs        string
	ExceptionSets        []string
	Telemetry            *string
	TelemetryV2          *string
	TelemetryV2Null      bool
	AnalyticSets         []AnalyticSetInput
	USBControlSet        *string
	CommsConfig          CommsConfigInput
	InfoSync             InfoSyncInput
	AutoUpdate           bool
	SignaturesFeedConfig SignaturesFeedConfigInput
}

// UpdatePlanRequest is the request payload for updating a plan
type UpdatePlanRequest struct {
	Name                 string
	Description          string
	LogLevel             *string
	ActionConfigs        string
	ExceptionSets        []string
	Telemetry            *string
	TelemetryV2          *string
	TelemetryV2Null      bool
	AnalyticSets         []AnalyticSetInput
	USBControlSet        *string
	CommsConfig          CommsConfigInput
	InfoSync             InfoSyncInput
	AutoUpdate           bool
	SignaturesFeedConfig SignaturesFeedConfigInput
}

// AnalyticSetInput is a plan analytic set input entry
type AnalyticSetInput struct {
	Type string
	UUID string
}

// CommsConfigInput captures communications configuration
type CommsConfigInput struct {
	FQDN     string
	Protocol string
}

// InfoSyncInput captures info sync configuration
type InfoSyncInput struct {
	Attrs                []string
	InsightsSyncInterval int64
}

// SignaturesFeedConfigInput captures signatures feed configuration
type SignaturesFeedConfigInput struct {
	Mode string
}

// ListPlansResponse represents the response from listing plans
type ListPlansResponse struct {
	Items    []Plan   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}

// PlanName is a lightweight plan containing only the name
type PlanName struct {
	Name string `json:"name"`
}

// ListPlanNamesResponse is the response wrapper for listing plan names
type ListPlanNamesResponse struct {
	Items    []PlanName `json:"items"`
	PageInfo PageInfo   `json:"pageInfo"`
}

// GetPlanConfigurationAndSetOptionsRequest holds RBAC flags for the plan configuration query
type GetPlanConfigurationAndSetOptionsRequest struct {
	RBACActionConfigs bool
	RBACTelemetry     bool
	RBACUSBControlSet bool
	RBACExceptionSet  bool
	RBACAnalyticSet   bool
}

// PlanConfigRefItem is a lightweight reference item returned by plan configuration queries
type PlanConfigRefItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PlanConfigExceptionSetItem is an exception set reference in plan configuration
type PlanConfigExceptionSetItem struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Managed bool   `json:"managed"`
}

// PlanConfigAnalyticSetItem is an analytic set reference in plan configuration
type PlanConfigAnalyticSetItem struct {
	Name        string   `json:"name"`
	UUID        string   `json:"uuid"`
	Description string   `json:"description"`
	Managed     bool     `json:"managed"`
	Types       []string `json:"types"`
}

// PlanConfigurationAndSetOptions is the combined response for plan configuration options
type PlanConfigurationAndSetOptions struct {
	ActionConfigs       []PlanConfigRefItem          `json:"actionConfigs"`
	Telemetries         []PlanConfigRefItem          `json:"telemetries"`
	TelemetriesV2       []PlanConfigRefItem          `json:"telemetriesV2"`
	USBControlSets      []PlanConfigRefItem          `json:"usbControlSets"`
	ExceptionSets       []PlanConfigExceptionSetItem `json:"exceptionSets"`
	AnalyticSets        []PlanConfigAnalyticSetItem  `json:"analyticSets"`
	ManagedAnalyticSets []PlanConfigAnalyticSetItem  `json:"managedAnalyticSets"`
}

// PlanConfigProfileOptionsInput controls which payloads are included in a plan configuration profile.
type PlanConfigProfileOptionsInput struct {
	PPPC              bool
	Token             bool
	TokenOptions      PlanConfigProfileTokenOptionsInput
	CA                bool
	CSR               bool
	Websocket         bool
	Sign              bool
	SystemExtension   bool
	ServiceManagement bool
	ConfigVersion     *int64
}

// PlanConfigProfileTokenOptionsInput controls bootstrap token payload options.
type PlanConfigProfileTokenOptionsInput struct {
	XPC              bool
	KeychainClientID bool
}
