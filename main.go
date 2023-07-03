package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState uint16

const (
	Running GameState = iota
	Exit
)

const (
	maxTicks = 60
)

func main() {

	if err := InitWindow(); err != nil {
		Logger.Panic(fmt.Sprint("Something went wrong starting up: ", err.Error()))
	}
	defer DestroyWindow()

	surface, err := Window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	Window.UpdateSurface()

	RunEngine()
}

func RunEngine() {
	var state GameState = Running
	lastFrameTime := time.Now()

	for state == Running {
		delta := time.Since(lastFrameTime).Seconds() * maxTicks
		lastFrameTime = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch /*evt :=*/ event.(type) {
			case *sdl.QuitEvent:
				state = Exit

				// case *sdl.MouseMotionEvent:
				// 	Logger.Debug(fmt.Sprint(evt.XRel, " ", evt.YRel))
			}
		}

		moo(delta)
	}
}

func moo(delta float64) {}
