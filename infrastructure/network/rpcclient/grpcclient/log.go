package grpcclient

import (
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
