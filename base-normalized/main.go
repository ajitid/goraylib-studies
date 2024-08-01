package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WinWidth  = 800
	WinHeight = 600
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(WinWidth, WinHeight, "base normalized")
	defer rl.CloseWindow()

	// Useful during development:
	// rl.SetWindowState(rl.FlagWindowUnfocused) // not supported in SDL backend atm
	// rl.SetWindowMonitor(1)

	// Create a camera
	camera := rl.Camera2D{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		// rl.ClearBackground(rl.Blank) // or
		rl.ClearBackground(rl.Black)

		// Set the camera offset to the center of the screen
		camera.Offset = rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
		// Set the zoom to normalize coordinates
		// This will make 1 unit equal to half the screen height
		camera.Zoom = float32(rl.GetScreenHeight()) / 2

		// anything drawn within this mode will in Normalized space
		rl.BeginMode2D(camera)

		// Draw a red circle at (0, 0), which will appear at the center of the screen
		rl.DrawCircle(0, 0, 1, rl.DarkPurple)

		// Draw a blue rectangle from (-0.5, -0.5) to (0.5, 0.5)
		// This will cover the central quarter of the screen
		rect := rl.Rectangle{
			-0.5, -0.5, // adjusted wrt to width and height as `DrawRectangleRec`'s anchor is still at top-left and not center
			1, 1,
		}
		rl.DrawRectangleRec(rect, rl.Purple)

		rl.EndMode2D()

		// rest of them will be drawn using the default (top-left pixel based) co-ordinate system
		/*
			Add ` | rl.FlagWindowHighdpi` to `rl.SetConfigFlags()` if this pink border is not drawn at the edges of the window
			Do note that in this case texture scaling should be done by you. See:
			- https://github.com/raysan5/raylib/issues/2566
			- https://github.com/raysan5/raylib/discussions/2999
		*/
		rl.DrawRectangleLines(0, 0, WinWidth, WinHeight, rl.Pink) // canvas border
		rl.DrawText("Normalized space", 6, 6, 28, rl.White)
		rl.DrawText("(0,0) at center", 6, 28+6, 20, rl.White) // 28+6 : previous text's fontSize + posY
		rl.DrawText("Y range is (-1,1), 1 is at bottom", 6, 28+6+20, 20, rl.White)
		rl.DrawText("X range is (-1,1), 1 is at right", 6, 28+6+20+20, 20, rl.White)

		rl.EndDrawing()
	}
}
