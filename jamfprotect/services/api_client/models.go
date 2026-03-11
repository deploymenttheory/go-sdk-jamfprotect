package apiclient

// ApiClient represents an API client in Jamf Protect.
type ApiClient struct {
	ClientID      string          `json:"clientId"`
	Name          string          `json:"name"`
	AssignedRoles []ApiClientRole `json:"assignedRoles"`
	Password      string          `json:"password"`
	Created       string          `json:"created"`
}

// ApiClientRole represents a role assigned to an API client.
type ApiClientRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateApiClientRequest is the request payload for creating an API client.
type CreateApiClientRequest struct {
	Name    string
	RoleIDs []string
}

// UpdateApiClientRequest is the request payload for updating an API client.
type UpdateApiClientRequest struct {
	Name    string
	RoleIDs []string
}

// ListApiClientsResponse represents the paginated response from listing API clients.
type ListApiClientsResponse struct {
	Items    []ApiClient `json:"items"`
	PageInfo PageInfo    `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
