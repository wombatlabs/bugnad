package rpcclient

import (
	"github.com/wombatlabs/bugnad/infrastructure/logger"
	"github.com/wombatlabs/bugnad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
