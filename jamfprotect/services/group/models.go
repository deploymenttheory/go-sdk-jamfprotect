package group

// Group represents a group in Jamf Protect.
type Group struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Connection    *GroupConnection `json:"connection"`
	AssignedRoles []GroupRole     `json:"assignedRoles"`
	AccessGroup   bool            `json:"accessGroup"`
	Created       string          `json:"created"`
	Updated       string          `json:"updated"`
}

// GroupConnection represents the identity provider connection associated with a group.
type GroupConnection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GroupRole represents a role assigned to a group.
type GroupRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateGroupRequest is the request payload for creating a group.
type CreateGroupRequest struct {
	Name         string
	ConnectionID *string
	AccessGroup  bool
	RoleIDs      []string
}

// UpdateGroupRequest is the request payload for updating a group.
type UpdateGroupRequest struct {
	Name        string
	AccessGroup bool
	RoleIDs     []string
}

// ListGroupsResponse represents the paginated response from listing groups.
type ListGroupsResponse struct {
	Items    []Group  `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
