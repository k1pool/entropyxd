package main

import (
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
