package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "5626f575fb4dfa60b8ef47caa046fe1d0e015423"

// The main version number that is being run at the moment.
const Version = "0.1.0"

const BuildDate = "2019-08-01-21:24:50"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
