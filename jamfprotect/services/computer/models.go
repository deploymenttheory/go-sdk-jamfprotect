package computer

// Computer represents a computer enrolled in Jamf Protect.
type Computer struct {
	UUID                    *string      `json:"uuid"`
	Serial                  *string      `json:"serial"`
	HostName                *string      `json:"hostName"`
	ModelName               *string      `json:"modelName"`
	OSMajor                 *int64       `json:"osMajor"`
	OSMinor                 *int64       `json:"osMinor"`
	OSPatch                 *int64       `json:"osPatch"`
	Arch                    *string      `json:"arch"`
	CertID                  *string      `json:"certid"`
	MemorySize              *int64       `json:"memorySize"`
	OSString                *string      `json:"osString"`
	KernelVersion           *string      `json:"kernelVersion"`
	InstallType             *string      `json:"installType"`
	Label                   *string      `json:"label"`
	Created                 *string      `json:"created"`
	Updated                 *string      `json:"updated"`
	Version                 *string      `json:"version"`
	Checkin                 *string      `json:"checkin"`
	ConfigHash              *string      `json:"configHash"`
	Tags                    []string     `json:"tags"`
	SignaturesVersion       *int64       `json:"signaturesVersion"`
	Plan                    *ComputerPlan `json:"plan"`
	InsightsStatsFail       *int64       `json:"insightsStatsFail"`
	InsightsUpdated         *string      `json:"insightsUpdated"`
	ConnectionStatus        *string      `json:"connectionStatus"`
	LastConnection          *string      `json:"lastConnection"`
	LastConnectionIP        *string      `json:"lastConnectionIp"`
	LastDisconnection       *string      `json:"lastDisconnection"`
	LastDisconnectionReason *string      `json:"lastDisconnectionReason"`
	WebProtectionActive     *bool        `json:"webProtectionActive"`
	FullDiskAccess          *string      `json:"fullDiskAccess"`
	PendingPlan             *int64       `json:"pendingPlan"`
}

// ComputerPlan represents a plan assigned to a computer.
type ComputerPlan struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
	Hash *string `json:"hash"`
}

// ListComputersResponse represents the paginated response from listing computers.
type ListComputersResponse struct {
	Items    []Computer `json:"items"`
	PageInfo PageInfo   `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
