{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type Version struct {
	BuildTime string `json:"build_time"`
	GitCommit string `json:"git_commit"`
	Name      string `json:"name"`
	UserAgent string `json:"user_agent"`
	Version   string `json:"version"`
}

func GetAppVersion() Version {
	return Version{
		BuildTime: conf.AppBuildTime(),
		GitCommit: conf.AppGitCommit(),
		Name:      conf.AppName(),
		UserAgent: conf.UserAgent(),
		Version:   conf.AppVersion(),
	}
}
