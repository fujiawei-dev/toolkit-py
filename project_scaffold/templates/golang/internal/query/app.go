{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type AppDescription struct {
	AppName    string `json:"app_name" example:"程序名: App"`
	AppVersion string `json:"app_version" example:"版本号: 1.0.0"`
	BuildTime  string `json:"build_time" example:"编译时间: 2022-01-19T09:04:33+08:00"`
	GitCommit  string `json:"git_commit" example:"最新提交: 76bf0375"`
	UserAgent  string `json:"user_agent" example:"用户代理: go/go1.17.1 os/windows"`
}

func GetAppDescription() AppDescription {
	return AppDescription{
		AppName:    conf.AppName(),
		AppVersion: conf.AppVersion(),
		BuildTime:  conf.AppBuildTime(),
		GitCommit:  conf.AppGitCommit(),
		UserAgent:  conf.UserAgent(),
	}
}
