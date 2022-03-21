{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/acl"
)

type UserResult struct {
	ID       uint     `json:"id" example:"1"` // 记录ID
	Username string   `json:"username" example:"用户名"`
	Role     acl.Role `json:"role" example:"角色"`
}
