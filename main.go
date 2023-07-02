package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"go.uber.org/zap"
)

const (
	maxTicks = 60
)

var (
	running bool
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	if err := InitWindow(); err != nil {
		logger.Panic(fmt.Sprint("Something went wrong starting up: ", err.Error()))
	}
	defer DestroyWindow()

	RunEngine()
}

func RunEngine() {
	running = true
	lastFrameTime := time.Now()

	for running {
		delta := time.Since(lastFrameTime).Seconds() * maxTicks
		lastFrameTime = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				return
			}
		}

		moo(delta)
	}
}

func moo(delta float64) {}
