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
