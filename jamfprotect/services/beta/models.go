package beta

// BetaName identifies a beta program in Jamf Protect.
type BetaName string

const (
	// BetaNameNGTP is the NGTP beta program identifier.
	BetaNameNGTP BetaName = "NGTP_BETA"
)

// BetaAcceptanceStatus represents a tenant's beta enrollment status.
type BetaAcceptanceStatus struct {
	BetaName          string `json:"betaName"`
	AcceptedTimestamp string `json:"acceptedTimestamp"`
	AcceptedUser      string `json:"acceptedUser"`
}

// AppInitializationData contains organization initialization data from the API.
type AppInitializationData struct {
	BetaAcceptanceStatus []BetaAcceptanceStatus `json:"betaAcceptanceStatus"`
}
