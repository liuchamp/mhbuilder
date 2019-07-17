package version

import (
	"fmt"
	"github.com/hashicorp/consul/version"
	"runtime"
	"time"
)

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit = version.GitCommit

// The main version number that is being run at the moment.
const Version = "0.1.0"

var BuildDate = time.Now().String()

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
