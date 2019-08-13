package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "28021c5e04e7292c87896e6c9a14d736eff05741"

// The main version number that is being run at the moment.
const Version = "0.1.2"

const BuildDate = "2019-08-08-16:15:51"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
