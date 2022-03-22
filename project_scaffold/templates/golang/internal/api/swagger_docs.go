{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// These structures defined below are only for generating Swagger API documentation, please do not use them elsewhere!

type Result struct{}

// @Summary      获取基本描述
// @Description  查询应用程序当前的版本号、编译时间、最新提交等基础信息
// @Tags         程序设置
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Success      200  {object}  query.Response{result=query.AppDescription}  "操作成功"
// @Router       /app/description [get]
func getAppDescription() {}

// @Summary      用户注册
// @Description  注册用户/注册账号/创建用户/创建账号
// @Tags         用户管理
// @Accept       json
// @Param        object  body  form.User  true  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /user [post]
func postUser() {}

// @Summary      修改用户
// @Description  修改用户角色、用户名等
// @Tags         用户管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int                      true   "ID"
// @Param        object  body  form.UserUpdate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /user/{id} [put]
func putUser() {}

// @Summary      修改用户密码
// @Description  修改用户密码
// @Tags         用户管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int              true   "ID"
// @Param        object  body  form.UserChangePassword  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /user/{id}/password [put]
func putUserPassword() {}

// @Summary      用户列表
// @Description  用户列表
// @Tags         用户管理
// @Accept       json
// @Param        page       query  int     false  "页码"    default(1)
// @Param        page_size  query  int     false  "每页数量"  Enums(10, 20)  default(10)
// @Param        username   query  string  false  "用户名"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.UserResult}}  "操作成功"
// @Router       /users [get]
func getUsers() {}

// @Summary      用户登录
// @Description  用户登录，获取 Authorization Token 参数
// @Tags         用户管理
// @Accept       json
// @Param        object  body  form.UserLogin  true  "参数"
// @Produce      json
// @Header       200  {string}  Authorization                  "鉴权"
// @Success      200  {object}  query.Response{result=string}  "result 即 Authorization 参数，该值同时存在于 Header 的 Authorization 字段中"
// @Router       /user/login [post]
func postUserLogin() {}

// @Summary      用户退出
// @Description  用户主动退出登录状态，当前 Authorization Token 参数将失效
// @Tags         用户管理
// @Accept       json
// @Param        Authorization  header  string  true  "鉴权，默认格式为 Bearer $token"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /user/logout [post]
func postUserLogout() {}

// @Summary      删除操作日志
// @Description  删除操作日志
// @Tags         操作日志
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /operation_log/{id} [delete]
func deleteOperationLog() {}

// @Summary      删除全部操作日志
// @Description  删除全部操作日志
// @Tags         操作日志
// @Accept       json
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /operation_logs [delete]
func deleteOperationLogs() {}

// @Summary      获取操作日志
// @Description  获取操作日志
// @Tags         操作日志
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.OperationLogResult  "操作成功"
// @Router       /operation_log/{id} [get]
func getOperationLog() {}

// @Summary      获取操作日志列表
// @Description  获取操作日志列表
// @Tags         操作日志
// @Accept       json
// @Security     ApiKeyAuth
// @Param        page        query  int     false  "页码"    default(1)
// @Param        page_size   query  int     false  "每页数量"  Enums(10, 20)  default(10)
// @Param        time_begin  query  string  false  "开始时间，比如 2021-10-01"
// @Param        time_end    query  string  false  "结束时间，比如 2022-10-01"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.OperationLogResult}}  "操作成功"
// @Router       /operation_logs [get]
func getOperationLogs() {}

// @Summary      获取区域数据
// @Description  获取区域数据
// @Tags         常量数据
// @Accept       application/x-www-form-urlencoded
// @Param        code  query  string  false  "区域代码"
// @Produce      json
// @Success      200  {object}  query.Response{result=query.RegionLocation}  "操作成功"
// @Router       /region/by [get]
func getRegionBy() {}
