package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
const GitCommit = "ad1824e0d8bf559f64bce3b9c1e3ee1cd32de2fb"

// The main version number that is being run at the moment.
const Version = "1.0.1"

const BuildDate = "2019-08-16-21:55:37"

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
