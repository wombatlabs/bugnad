package connmanager

import (
	"github.com/wombatlabs/bugnad/infrastructure/logger"
	"github.com/wombatlabs/bugnad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
