{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

var Permissions = ACL{
	ResourceDefault: Roles{
		RoleAdmin:   Actions{ActionDefault: true},
		RoleDefault: Actions{ActionDefault: true},
		RoleGuest:   Actions{ActionDefault: false},
	},
}
