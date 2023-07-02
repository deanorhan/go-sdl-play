package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"modernc.org/mathutil"
)

const (
	minScreenWidth  = 800
	maxScreenWidth  = 2560
	minScreenHeight = 600
	maxScreenHeight = 1440
)

var (
	Window *sdl.Window

	cfg *Config
)

func init() {
	runtime.LockOSThread()
}

func loadConfig() error {
	var err error
	cfg, err = NewConfig("./engine.yml")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err.Error())
	}

	return nil
}

func InitWindow() error {
	Logger.Debug("Starting up the engine")

	if err := loadConfig(); err != nil {
		return err
	}

	Logger.Debug(fmt.Sprint("We have a config now, debug is ", cfg.Debug))

	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("SDL failed to init: %v", err.Error())
	}

	Logger.Debug("SDL initialized")

	screenWidth := mathutil.ClampInt32(cfg.Display.Width, minScreenWidth, maxScreenWidth)
	screenHeight := mathutil.ClampInt32(cfg.Display.Height, minScreenHeight, maxScreenHeight)

	Window, err = sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	Logger.Debug("Window created")

	return err
}

func DestroyWindow() {
	Logger.Debug("Shutting down the engine")

	if Window != nil {
		Window.Destroy()
		Logger.Debug("Window destroyed")
	}

	sdl.Quit()
	Logger.Debug("SDL shutdown")
	Logger.Sync()
}
