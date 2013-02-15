package main

import (
	"flag"
	"fmt"
	"github.com/go-gl/glfw"
	"github.com/wibbe/glh/gfx"
	"os"
)

var connectionAddress = flag.String("connect", "localhost", "Server IP address")
var levelToLoad = flag.String("level", "", "Level to load (Only use if you start a server)")

func main() {
	fmt.Printf("Little Shooter\n")

	flag.Parse()

	if *connectionAddress == "localhost" && *levelToLoad == "" {
		fmt.Fprintf(os.Stderr, "Must specify a level to load when running a server")
		return
	}

	glfw.Init()
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(800, 550, 8, 8, 8, 8, 24, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetWindowTitle("Little Shooter")

	ctx, err := gfx.NewGL32Context()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	game := NewGame(*connectionAddress)

	glfw.SetTime(0.0)
	simulationTime := 0.0

	continueOn := true
	for continueOn {
		gameTime := glfw.Time()

		for simulationTime < gameTime {
			game.Tick(FrameTick)
			simulationTime += FrameTick
		}

		ctx.Clear(true, true, true)
		glfw.SwapBuffers()

		continueOn = glfw.WindowParam(glfw.Opened) == 1 && glfw.Key(glfw.KeyEsc) == glfw.KeyRelease
	}
}
