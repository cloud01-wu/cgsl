package role

// Role defines permission of users
type Role string

type IRole interface{}

const (
	// RoleUnknown is invalid role
	RoleUnknown = ""
	// RoleAdmin has permission to access all pages
	RoleAdmin = "admin"
	// RoleEditor has permission to see all pages except permission page
	RoleEditor = "editor"
	// RoleVisitor can only see dashboard and document pages
	RoleVisitor = "visitor"
)

func (e Role) String() string {
	return string(e)
}

func FindRole(role string) Role {
	switch role {
	case "admin":
		return RoleAdmin
	case "editor":
		return RoleEditor
	case "visitor":
		return RoleVisitor
	default:
		return RoleUnknown
	}
}
