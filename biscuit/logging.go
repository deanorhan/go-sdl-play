package biscuit

import (
	"os"

	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	debug  bool
)

func init() {
	debug = true
	var loggerConfig zap.Config

	if os.Getenv("GOENV") == "prod" {
		loggerConfig = zap.NewProductionConfig()
		debug = false

	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	Logger, _ = loggerConfig.Build()
}

func EnableDebug(enableDebug bool) {
	if debug && !enableDebug {
		Logger, _ = zap.NewProduction()

	} else if !debug && enableDebug {
		Logger, _ = zap.NewDevelopment()
	}

	debug = enableDebug
}
