package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
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
	Window  *sdl.Window
	Context sdl.GLContext

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

	if err != nil {
		return fmt.Errorf("SDL failed to create window: %v", err.Error())
	}

	Logger.Debug("Window created")

	Context, err = Window.GLCreateContext()
	if err != nil {
		panic(err)
	}

	Logger.Debug("Created OpenGL context")

	initOpenGL()
	return err
}

func DestroyWindow() {
	Logger.Debug("Shutting down the engine")

	if Context != nil {
		sdl.GLDeleteContext(Context)
		Logger.Debug("Context destroyed")
	}

	if Window != nil {
		Window.Destroy()
		Logger.Debug("Window destroyed")
	}

	sdl.Quit()
	Logger.Debug("SDL shutdown")
	Logger.Sync()
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	Logger.Debug(fmt.Sprint("OpenGL version", version))

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}
