package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "61e5164764594dd24c4ab4376bcc9520c1f6d599"

// The main version number that is being run at the moment.
const Version = "0.1.1"

const BuildDate = "2019-08-07-15:01:41"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
