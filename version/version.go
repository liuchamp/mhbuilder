package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "c4480db175c1b83efcbc10a4f3c4d87796ba489d"

// The main version number that is being run at the moment.
const Version = "1.0.1"

const BuildDate = "2019-08-16-21:44:11"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
