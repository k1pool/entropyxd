package main

import (
	"fmt"
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
	"os"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("KSMN")
	spawn      = panics.GoroutineWrapperFunc(log)
)

func initLog(logFile, errLogFile string) {
	log.SetLevel(logger.LevelDebug)
	err := backendLog.AddLogFile(logFile, logger.LevelTrace)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding log file %s as log rotator for level %s: %s", logFile, logger.LevelTrace, err)
		os.Exit(1)
	}
	err = backendLog.AddLogFile(errLogFile, logger.LevelWarn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding log file %s as log rotator for level %s: %s", errLogFile, logger.LevelWarn, err)
		os.Exit(1)
	}
	err = backendLog.AddLogWriter(os.Stdout, logger.LevelInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding stdout to the loggerfor level %s: %s", logger.LevelWarn, err)
		os.Exit(1)
	}
	err = backendLog.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting the logger: %s ", err)
		os.Exit(1)
	}

}
