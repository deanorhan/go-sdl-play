package render

import (
	"fmt"

	"github.com/deanorhan/go-sdl-play/util"
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

	cfg *util.Config
)

func loadConfig() error {
	var err error
	cfg, err = util.NewConfig("./engine.yml")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err.Error())
	}

	return nil
}

func InitWindow() error {
	util.Logger.Debug("Starting up the engine")

	if err := loadConfig(); err != nil {
		return err
	}

	util.Logger.Debug(fmt.Sprint("We have a config now, debug is ", cfg.Debug))

	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("SDL failed to init: %v", err.Error())
	}

	util.Logger.Debug("SDL initialized")

	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG, 1)

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

	util.Logger.Debug("Window created")

	Context, err = Window.GLCreateContext()
	if err != nil {
		return fmt.Errorf("GL Context creation failed: %v", err.Error())
	}

	util.Logger.Debug("Created OpenGL context")

	return initOpenGL()
}

func DestroyWindow() {
	util.Logger.Debug("Shutting down the engine")

	if Context != nil {
		sdl.GLDeleteContext(Context)
		util.Logger.Debug("Context destroyed")
	}

	if Window != nil {
		Window.Destroy()
		util.Logger.Debug("Window destroyed")
	}

	sdl.Quit()
	util.Logger.Debug("SDL shutdown")
	util.Logger.Sync()
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() error {
	if err := gl.Init(); err != nil {
		return fmt.Errorf("GL failed to init: %v", err.Error())
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	util.Logger.Debug(fmt.Sprint("OpenGL version", version))

	gl.Viewport(0, 0, cfg.Display.Width, cfg.Display.Height)
	return nil
}
