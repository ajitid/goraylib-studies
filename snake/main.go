// Following https://www.youtube.com/watch?v=lfiQNCNUifI
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WinSize        = 600
	GridWidth      = 20
	CellSize       = 16
	CanvasSize     = GridWidth * CellSize
	TickRate       = 0.13
	MaxSnakeLength = GridWidth * GridWidth
)

var (
	tickTimer     float32 = TickRate
	snake         [MaxSnakeLength]rl.Vector2
	snakeLength   int32
	moveDirection rl.Vector2
	gameOver      bool
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(WinSize, WinSize, "snake")
	defer rl.CloseWindow()

	rl.SetWindowMonitor(1)

	restart()

	for !rl.WindowShouldClose() {
		update()

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

func restart() {
	snakeHeadPos := rl.Vector2{GridWidth / 2, GridWidth / 2}
	snake[0], snake[1], snake[2] = snakeHeadPos, snakeHeadPos, snakeHeadPos
	snake[1].X -= 1
	snake[2].X -= 2
	snakeLength = 3

	moveDirection = rl.Vector2{1, 0}

	gameOver = false
}

func update() {
	switch {
	case rl.IsKeyPressed(rl.KeyUp):
		if !(moveDirection.X == 0 && moveDirection.Y == 1) { // If already moving down, don't let it move upwards. Otherwise snake will backtrack and will move on itself
			moveDirection.X, moveDirection.Y = 0, -1
		}
	case rl.IsKeyPressed(rl.KeyDown):
		if !(moveDirection.X == 0 && moveDirection.Y == -1) {
			moveDirection.X, moveDirection.Y = 0, 1
		}
	case rl.IsKeyPressed(rl.KeyRight):
		if !(moveDirection.X == -1 && moveDirection.Y == 0) {
			moveDirection.X, moveDirection.Y = 1, 0
		}
	case rl.IsKeyPressed(rl.KeyLeft):
		if !(moveDirection.X == 1 && moveDirection.Y == 0) {
			moveDirection.X, moveDirection.Y = -1, 0
		}
	}

	if gameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			restart()
		}
	} else {
		tickTimer -= rl.GetFrameTime()
	}

	if tickTimer <= 0 {
		nextPartPos := snake[0]
		snake[0].X += moveDirection.X
		snake[0].Y += moveDirection.Y

		headPos := snake[0]
		if headPos.X < 0 || headPos.Y < 0 || headPos.X >= GridWidth || headPos.Y >= GridWidth {
			gameOver = true
		}

		for i := int32(1); i < snakeLength; i++ {
			snake[i], nextPartPos = nextPartPos, snake[i]
		}
		tickTimer = TickRate + tickTimer
	}
}

func draw() {
	for i := range snakeLength {
		headRect := rl.Rectangle{
			snake[i].X * CellSize,
			snake[i].Y * CellSize,
			CellSize,
			CellSize,
		}
		rl.DrawRectangleRec(headRect, rl.White)
	}

	if gameOver {
		rl.DrawText("Game Over!", 4, 4, 25, rl.Red)
		rl.DrawText("Press Enter to play again!", 4, 30, 15, rl.Black)
	}
}
