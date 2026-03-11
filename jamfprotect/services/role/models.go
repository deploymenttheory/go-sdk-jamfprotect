package role

// Role represents a role in Jamf Protect.
type Role struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Permissions RolePermissions `json:"permissions"`
	Created     string          `json:"created"`
	Updated     string          `json:"updated"`
}

// RolePermissions represents the read and write permissions for a role.
type RolePermissions struct {
	Read  []string `json:"R"`
	Write []string `json:"W"`
}

// CreateRoleRequest is the request payload for creating a role.
type CreateRoleRequest struct {
	Name           string
	ReadResources  []string
	WriteResources []string
}

// UpdateRoleRequest is the request payload for updating a role.
type UpdateRoleRequest struct {
	Name           string
	ReadResources  []string
	WriteResources []string
}

// ListRolesResponse represents the paginated response from listing roles.
type ListRolesResponse struct {
	Items    []Role   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
