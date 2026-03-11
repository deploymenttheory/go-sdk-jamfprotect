package user

// User represents a user in Jamf Protect.
type User struct {
	ID                    string       `json:"id"`
	Email                 string       `json:"email"`
	Sub                   *string      `json:"sub"`
	Connection            *UserConnection `json:"connection"`
	AssignedRoles         []UserRole   `json:"assignedRoles"`
	AssignedGroups        []UserGroup  `json:"assignedGroups"`
	LastLogin             *string      `json:"lastLogin"`
	Source                string       `json:"source"`
	ReceiveEmailAlert     bool         `json:"receiveEmailAlert"`
	EmailAlertMinSeverity string       `json:"emailAlertMinSeverity"`
	ExtraAttrs            string       `json:"extraAttrs"`
	Created               string       `json:"created"`
	Updated               string       `json:"updated"`
}

// UserConnection represents the identity provider connection associated with a user.
type UserConnection struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	RequireKnownUsers bool   `json:"requireKnownUsers"`
	Source            string `json:"source"`
}

// UserRole represents a role assigned to a user.
type UserRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserGroup represents a group assigned to a user.
type UserGroup struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	AssignedRoles []UserRole `json:"assignedRoles"`
}

// CreateUserRequest is the request payload for creating a user.
type CreateUserRequest struct {
	Email                 string
	RoleIDs               []string
	GroupIDs              []string
	ConnectionID          *string
	ReceiveEmailAlert     bool
	EmailAlertMinSeverity string
}

// UpdateUserRequest is the request payload for updating a user.
type UpdateUserRequest struct {
	RoleIDs               []string
	GroupIDs              []string
	ReceiveEmailAlert     bool
	EmailAlertMinSeverity string
}

// ListUsersResponse represents the paginated response from listing users.
type ListUsersResponse struct {
	Items    []User   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information.
type PageInfo struct {
	Next  *string `json:"next"`
	Total int     `json:"total"`
}
