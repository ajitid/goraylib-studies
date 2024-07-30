// Following https://www.youtube.com/watch?v=lfiQNCNUifI
package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WinSize    = 600
	GridWidth  = 20
	CellSize   = 16
	CanvasSize = GridWidth * CellSize
	TickRate   = 0.13
)

var tickTimer float32 = TickRate
var snakeHeadPos = rl.Vector2{GridWidth / 2, GridWidth / 2}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(WinSize, WinSize, "snake")
	defer rl.CloseWindow()

	rl.SetWindowMonitor(1)

	for !rl.WindowShouldClose() {
		tickTimer -= rl.GetFrameTime()
		if tickTimer <= 0 {
			snakeHeadPos.Y += 1
			tickTimer = TickRate + tickTimer
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(76, 53, 83, 255))

		camera := rl.Camera2D{
			Zoom: float32(WinSize) / CanvasSize,
		}
		rl.BeginMode2D(camera)
		draw()
		rl.EndMode2D()

		rl.EndDrawing()
	}
}

func draw() {
	headRect := rl.Rectangle{
		snakeHeadPos.X * CellSize,
		snakeHeadPos.Y * CellSize,
		CellSize,
		CellSize,
	}
	rl.DrawRectangleRec(headRect, rl.White)
}
