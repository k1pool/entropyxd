package main

import (
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)
