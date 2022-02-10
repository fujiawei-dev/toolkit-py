{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/iris-contrib/swagger/v12"
	irisSwagger "github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/docs"
)

func RegisterSwagger(router iris.Party) {
	if !conf.Swagger() {
		return
	}

	docs.SwaggerInfo.Title = conf.AppTitle()
	docs.SwaggerInfo.Description = conf.AppDescription()
	docs.SwaggerInfo.Host = conf.ExternalHttpHostPort()
	docs.SwaggerInfo.BasePath = conf.BasePath()
	docs.SwaggerInfo.Version = conf.AppVersion()

	swaggerUI := swagger.WrapHandler(irisSwagger.Handler)

	// Register on /swagger
	router.Get("/swagger", swaggerUI)

	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	router.Get("/swagger/{any:path}", swaggerUI)

	log.Infof("swagger: http://%s/swagger/index.html", conf.ExternalHttpHostPort()+conf.BasePath())
}

// 以下定义的数据结构仅是为了生成接口文档，请勿用在其他地方！

type httpResponseBody struct {
	Code   int    `json:"code" example:"0"`     // 错误码，正常情况 0
	Msg    string `json:"message" example:"OK"` // 响应消息
	Err    string `json:"error" example:"错误详情，仅测试模式可见，前端可忽略"`
	Result string `json:"result" example:"请求结果"`
}

type pager struct {
	Page      int `json:"page" form:"page,default=1"`                                    // 页码
	PageSize  int `json:"page_size" form:"page_size,default=10" binding:"gte=5,lte=100"` // 每页数量
	TotalRows int `json:"total_rows"`                                                    // 总行数
}

type userRequestBody struct {
	UserName string `json:"username" binding:"required" example:"用户名"`
	Password string `json:"password" binding:"required" example:"密码"`
}

type userLoginResponseBody struct {
	httpResponseBody
	Result string `json:"result" example:"Authorization 参数，该值同时存在于 Header 的 Authorization 中"`
}

type userLoginRefreshResponseBody struct {
	httpResponseBody
	Result struct {
		AccessToken  string `json:"access_token" example:"Authorization 参数"`
		RefreshToken string `json:"refresh_token" example:"refresh_token 参数，可用于刷新登录状态"`
	}
}

type versionResponseBody struct {
	httpResponseBody
	Result struct {
		BuildTime string `json:"build_time" example:"编译时间: 2022-01-19T09:04:33+08:00"`
		GitCommit string `json:"git_commit" example:"最新提交: 76bf0375"`
		Name      string `json:"name" example:"程序名: App"`
		UserAgent string `json:"user_agent" example:"用户代理: go/go1.17.1 os/windows"`
		Version   string `json:"version" example:"版本号: 1.0.0"`
	}
}

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户管理
// @Accept json
// @Param object body userRequestBody true "用户注册信息"
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /user [post]
func postUser(c iris.Context) {}

// @Summary 用户登录
// @Description 用户登录，获取 Authorization Token 参数，不带 Bearer 等前缀，前端请自行添加，有效期默认为 3 小时，可刷新状态模式下则一次登录有效期为 3x8 小时，但单个 Token 有效期仍为 3 小时
// @Tags 用户管理
// @Accept json
// @Param object body userRequestBody true "用户登录信息"
// @Produce json
// @Success 200 {object} userLoginResponseBody "默认模式"
// @Header 200 {string} Authorization "鉴权"
// @Success 201 {object} userLoginRefreshResponseBody "可刷新模式"
// @Router /user/login [post]
func postUserLogin(c iris.Context) {}

// @Summary 用户登录状态刷新
// @Description 用户登录状态刷新，这种模式下允许前端延长单次登录的有效期，但单个 Token 有效期不变
// @Tags 用户管理
// @Param Authorization header string true "鉴权，默认格式为 Bearer $token"
// @Param q query string true "登录接口或者该接口返回的 refresh_token 值"
// @Produce json
// @Success 200 {object} userLoginRefreshResponseBody "操作成功"
// @Header 200 {string} Authorization "鉴权，不带 Bearer 等前缀，前端请自行添加"
// @Router /user/login/refresh [get]
func postUserLoginRefresh(c iris.Context) {}

// @Summary 用户退出
// @Description 用户主动退出登录状态，当前 Authorization 值将失效
// @Tags 用户管理
// @Accept json
// @Param Authorization header string true "鉴权，默认格式为 Bearer $token"
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /user/logout [post]
func postUserLogout(c iris.Context) {}

// @Summary 查询版本
// @Description 查询程序当前的版本号、编译时间、最新提交等信息
// @Tags 系统管理
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /system/version [get]
func getVersion(c iris.Context) {}
