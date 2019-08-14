package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "d38a81e4b870cde93b6dcb14daf9f218865437be"

// The main version number that is being run at the moment.
const Version = "1.0.0"

const BuildDate = "2019-08-13-14:38:37"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
