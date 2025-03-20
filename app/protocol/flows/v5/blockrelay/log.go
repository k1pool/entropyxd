package blockrelay

import (
	"github.com/k1pool/entropyxd/infrastructure/logger"
	"github.com/k1pool/entropyxd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
