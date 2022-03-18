{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// Action 动作/操作
type (
	Action  string
	Actions map[Action]bool
)

const (
	ActionDefault Action = "default" // 默认
	ActionLogin   Action = "login"   // 登录
	ActionLogout  Action = "logout"  // 登出
	ActionSearch  Action = "search"  // 搜索
	ActionCreate  Action = "create"  // 创建
	ActionRead    Action = "read"    // 读取
	ActionUpdate  Action = "update"  // 修改
	ActionDelete  Action = "delete"  // 删除
)
