{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type Action string
type Actions map[Action]bool

const (
	ActionDefault Action = "default"
	ActionSearch  Action = "search"
	ActionCreate  Action = "create"
	ActionRead    Action = "read"
	ActionUpdate  Action = "update"
	ActionDelete  Action = "delete"
)
