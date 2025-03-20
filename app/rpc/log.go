package rpc

import (
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
