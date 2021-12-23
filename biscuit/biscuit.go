package biscuit

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"modernc.org/mathutil"
)

var (
	Window *sdl.Window
	// renderer *sdl.Renderer
	running bool

	cfg *Config
)

const (
	minScreenWidth  = 800
	maxScreenWidth  = 2560
	minScreenHeight = 600
	maxScreenHeight = 1440
)

func loadConfig() error {
	var err error
	cfg, err = NewConfig("./engine.yml")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err.Error())
	}

	return nil
}

func InitEngine() error {
	Logger.Debug("Starting up the engine")

	if err := loadConfig(); err != nil {
		return err
	}
	Logger.Debug(fmt.Sprint("We have a config now, debug is ", cfg.Debug))

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("SDL failed to init: %v", err.Error())
	}

	Logger.Debug("SDL initialized")

	if err := createWindow(); err != nil {
		return fmt.Errorf("SDL Window failed to init: %v", err.Error())
	}

	Logger.Debug("Window created")

	return nil
}

func createWindow() error {
	screenWidth := mathutil.ClampInt32(cfg.Display.Width, minScreenWidth, maxScreenWidth)
	screenHeight := mathutil.ClampInt32(cfg.Display.Height, minScreenHeight, maxScreenHeight)

	var err error
	Window, err = sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	return err
}

func RunEngine() {
	running = true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				return
			}
		}
	}
}

func StopEngine() {
	running = false
}

func ShutdownEngine() {
	Logger.Debug("Shutting down the engine")

	if Window != nil {
		Window.Destroy()
		Logger.Debug("Window destroyed")
	}

	sdl.Quit()
	Logger.Debug("SDL shutdown")
	Logger.Sync()
}
