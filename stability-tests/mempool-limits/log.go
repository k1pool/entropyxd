package mempoollimits

import (
	"fmt"
	"os"

	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/stability-tests/common"
	"github.com/k1pool/entropyxd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MPLM")
	spawn      = panics.GoroutineWrapperFunc(log)
)

func initLog(logFile, errLogFile string) {
	level := logger.LevelInfo
	if activeConfig().LogLevel != "" {
		var ok bool
		level, ok = logger.LevelFromString(activeConfig().LogLevel)
		if !ok {
			fmt.Fprintf(os.Stderr, "Log level %s doesn't exists", activeConfig().LogLevel)
			os.Exit(1)
		}
	}
	log.SetLevel(level)
	common.InitBackend(backendLog, logFile, errLogFile)
}
