package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "ed065dc4b469221a871460b6c8960618d7985dac"

// The main version number that is being run at the moment.
const Version = "0.1.1"

const BuildDate = "2019-08-01-21:30:55"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
