package main

import (
	"github.com/wombatlabs/bugnad/infrastructure/logger"
	"github.com/wombatlabs/bugnad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MNJS")
	spawn      = panics.GoroutineWrapperFunc(log)
)
