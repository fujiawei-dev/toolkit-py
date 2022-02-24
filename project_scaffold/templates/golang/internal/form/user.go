{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// User required information for creation
type User struct {
	Username string `json:"username" validate:"required" binding:"required" example:"用户名"`
	Password string `json:"password" validate:"required" binding:"required" example:"密码"`
}

type UserLogin struct {
	Username string `json:"username" form:"username" validate:"required" binding:"required" example:"用户名"`
	Password string `json:"password" form:"password" validate:"required" binding:"required" example:"密码"`
}

type UserChangePassword struct {
	OldPassword string `json:"old_password" example:"旧密码"`
	NewPassword string `json:"new_password" validate:"required" binding:"required" example:"新密码"`
}
