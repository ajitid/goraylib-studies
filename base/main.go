package main

import (
	"goraylib-studies/ctrl"
	"goraylib-studies/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	winWidth  int32 = 800
	winHeight int32 = 600
)

func main() {
	go ctrl.RunServer()
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(winWidth, winHeight, "base")
	defer rl.CloseWindow()

	// Useful during development:
	// rl.SetWindowState(rl.FlagWindowUnfocused) // not supported in SDL backend atm
	rl.SetWindowMonitor(1)
	go util.FocusEditor()

	for !rl.WindowShouldClose() {
		winWidth = int32(rl.GetScreenWidth())
		winHeight = int32(rl.GetScreenHeight())

		rl.BeginDrawing()
		ctrl.IsDrawing = true
		// rl.ClearBackground(rl.Blank) // or
		rl.ClearBackground(rl.NewColor(76, 53, 83, 255))

		rl.DrawRectangle(60, 60, 60, 30, rl.RayWhite) // random positioned rect
		rl.DrawRectangle(0, 0, 14, 14, rl.Blue)       // top-left

		/*
			Add ` | rl.FlagWindowHighdpi` to `rl.SetConfigFlags()` if this pink border is not drawn at the edges of the window
			Do note that in this case texture scaling should be done by you. See:
			- https://github.com/raysan5/raylib/issues/2566
			- https://github.com/raysan5/raylib/discussions/2999
		*/
		rl.DrawRectangleLines(0, 0, winWidth, winHeight, rl.Pink) // canvas border

		ctrl.IsDrawing = false
		rl.EndDrawing()
	}
}
