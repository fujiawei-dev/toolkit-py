{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type Permission struct {
	Roles   Roles
	Actions Actions
}

type ACL map[Resource]Roles

func (l ACL) Deny(resource Resource, role Role, action Action) bool {
	return !l.Allow(resource, role, action)
}

func (l ACL) Allow(resource Resource, role Role, action Action) bool {
	if p, ok := l[resource]; ok {
		return p.Allow(role, action)
	} else if p, ok = l[ResourceDefault]; ok {
		return p.Allow(role, action)
	}

	return false
}

func (a Actions) Allow(action Action) bool {
	if result, ok := a[action]; ok {
		return result
	} else if result, ok = a[ActionDefault]; ok {
		return result
	}

	return false
}

func (r Roles) Allow(role Role, action Action) bool {
	if a, ok := r[role]; ok {
		return a.Allow(action)
	} else if a, ok = r[RoleDefault]; ok {
		return a.Allow(action)
	}

	return false
}
