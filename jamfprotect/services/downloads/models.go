package downloads

// OrganizationDownloads represents the download resources available for a Jamf Protect organization.
type OrganizationDownloads struct {
	PPPC                    string         `json:"pppc"`
	RootCA                  string         `json:"rootCA"`
	CSR                     string         `json:"csr"`
	InstallerUUID           string         `json:"installerUuid"`
	VanillaPackage          *VanillaPackage `json:"vanillaPackage"`
	WebsocketAuth           string         `json:"websocket_auth"`
	TamperPreventionProfile string         `json:"tamperPreventionProfile"`
}

// VanillaPackage represents a vanilla package with version information.
type VanillaPackage struct {
	Version string `json:"version"`
}
