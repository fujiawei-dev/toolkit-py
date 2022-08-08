package acl

var Permissions = ACL{
	ResourceDefault: Roles{
		RoleAdmin:   Actions{ActionDefault: true},
		RoleDefault: Actions{ActionDefault: true},
		RoleGuest:   Actions{ActionDefault: false},
	},
}
