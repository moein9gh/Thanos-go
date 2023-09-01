package build_info

import (
	"fmt"
	"runtime"

	"github.com/thanos-go/log"
)

// Following variables are set via -ldflags
var (
	// AppVersion git SHA at build time
	AppVersion string
	// BuildTime time of build
	BuildTime string
	// VCSRef name of branch at build time
	VCSRef string
)

func Print() {
	log.Info("build info", map[string]interface{}{
		"app_version": AppVersion,
		"go_version":  runtime.Version(),
		"git_commit":  VCSRef,
		"build_time":  BuildTime,
		"os/arch":     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}
