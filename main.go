package main

import (
	"fmt"

	"github.com/deanorhan/biscuit"
	"github.com/deanorhan/biscuit/ecs"
	"go.uber.org/zap"
)

type DebugSystem struct {
}

func (ds *DebugSystem) Init(*ecs.World)       {}
func (ds *DebugSystem) Process(delta float64) { fmt.Println(delta) }

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	if err := biscuit.InitEngine(); err != nil {
		logger.Panic(fmt.Sprint("Something went wrong starting up: ", err.Error()))
	}
	defer biscuit.ShutdownEngine()

	world := biscuit.NewWorld()
	world.NewEntity()

	// biscuit.GetWorld().AddSystem(&DebugSystem{})

	biscuit.RunEngine()
}
