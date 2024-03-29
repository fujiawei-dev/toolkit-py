package version

import (
	"context"
	"fmt"
	"runtime"

	"github.com/docker/docker/pkg/parsers/kernel"
	"github.com/docker/docker/pkg/useragent"
)

// UAStringKey is used as key type for user-agent string in net/context struct
type UAStringKey struct{}

// UserAgent is the User-Agent the App client uses to identify itself.
// In accordance with RFC 7231 (5.5.3) is of the form:
//    [app client's UA] UpstreamClient([upstream client's UA])
func UserAgent(ctx context.Context) string {
	httpVersion := make([]useragent.VersionInfo, 0, 6)
	httpVersion = append(httpVersion, useragent.VersionInfo{Name: Name, Version: Version})
	httpVersion = append(httpVersion, useragent.VersionInfo{Name: "go", Version: runtime.Version()})
	httpVersion = append(httpVersion, useragent.VersionInfo{Name: "git-commit", Version: GitCommit})
	if kernelVersion, err := kernel.GetKernelVersion(); err == nil {
		httpVersion = append(httpVersion, useragent.VersionInfo{Name: "kernel", Version: kernelVersion.String()})
	}
	httpVersion = append(httpVersion, useragent.VersionInfo{Name: "os", Version: runtime.GOOS})
	httpVersion = append(httpVersion, useragent.VersionInfo{Name: "arch", Version: runtime.GOARCH})

	appUA := useragent.AppendVersions("", httpVersion...)
	upstreamUA := getUserAgentFromContext(ctx)
	if len(upstreamUA) > 0 {
		ret := insertUpstreamUserAgent(upstreamUA, appUA)
		return ret
	}
	return appUA
}

// getUserAgentFromContext returns the previously saved user-agent context stored in ctx, if one exists
func getUserAgentFromContext(ctx context.Context) string {
	var upstreamUA string
	if ctx != nil {
		var ki = ctx.Value(UAStringKey{})
		if ki != nil {
			upstreamUA = ctx.Value(UAStringKey{}).(string)
		}
	}
	return upstreamUA
}

// escapeStr returns s with every rune in charsToEscape escaped by a backslash
func escapeStr(s string, charsToEscape string) string {
	var ret string
	for _, currRune := range s {
		appended := false
		for _, escapableRune := range charsToEscape {
			if currRune == escapableRune {
				ret += `\` + string(currRune)
				appended = true
				break
			}
		}
		if !appended {
			ret += string(currRune)
		}
	}
	return ret
}

// insertUpstreamUserAgent adds the upstream client useragent to create a user-agent
// string of the form:
//   $appUA UpstreamClient($upstreamUA)
func insertUpstreamUserAgent(upstreamUA string, appUA string) string {
	charsToEscape := `();\`
	upstreamUAEscaped := escapeStr(upstreamUA, charsToEscape)
	return fmt.Sprintf("%s UpstreamClient(%s)", appUA, upstreamUAEscaped)
}
