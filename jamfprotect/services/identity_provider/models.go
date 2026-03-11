package identityprovider

// Connection represents an identity provider connection in Jamf Protect.
type Connection struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	RequireKnownUsers bool   `json:"requireKnownUsers"`
	Button            string `json:"button"`
	Created           string `json:"created"`
	Updated           string `json:"updated"`
	Strategy          string `json:"strategy"`
	GroupsSupport     bool   `json:"groupsSupport"`
	Source            string `json:"source"`
}

// ListConnectionsResponse represents the paginated response from listing connections.
type ListConnectionsResponse struct {
	Items    []Connection `json:"items"`
	PageInfo PageInfo     `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
