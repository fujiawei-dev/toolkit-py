{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type Role string
type Roles map[Role]Actions

const (
	RoleAdmin   Role = "admin"
	RoleDefault Role = "default"
	RoleGuest   Role = "guest"
)
