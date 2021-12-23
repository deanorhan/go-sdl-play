package main

import (
	"fmt"

	"github.com/deanorhan/biscuit"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	if err := biscuit.InitEngine(); err != nil {
		logger.Panic(fmt.Sprint("Something went wrong starting up: ", err.Error()))
	}
	defer biscuit.ShutdownEngine()

	biscuit.RunEngine()
}
