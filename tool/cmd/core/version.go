package core

import (
	"fmt"
	"runtime"
	"time"
)

// const vars related to the app version
const (
	MajorVersion = "1"
	MinorVersion = "0"

	PrereleaseBlurb = "This application is inteded to be used as a tool to developers."
	IsRelease       = false
	GitHub          = "Github: https://github.com/leonel-garofolo/sodaProject"
)

// vars related to the app version
var (
	Copyright = fmt.Sprintf("Copyright (c) 2021-%d The Oesia Developers.",
		time.Now().Year())
)

// Version returns the version string
func Version(short bool) string {
	versionStr := fmt.Sprintf("Soda v%s.%s %s %s",
		MajorVersion, MinorVersion, runtime.GOARCH, runtime.Version())
	if !IsRelease {
		versionStr += " pre-release.\n"
		if !short {
			versionStr += PrereleaseBlurb + "\n"
		}
	} else {
		versionStr += " release.\n"
	}
	if short {
		return versionStr
	}
	versionStr += Copyright + "\n\n"
	versionStr += GitHub
	return versionStr
}
