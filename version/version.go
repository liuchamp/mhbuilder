package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "0f42b741d8bae5e77c764fdc086c81b9ce46d573"

// The main version number that is being run at the moment.
const Version = "0.1.1"

const BuildDate = "2019-08-08-16:11:44"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
