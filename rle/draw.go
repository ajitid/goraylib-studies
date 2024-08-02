package rle

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Draw a color-filled rectangle with anchor at its center
func DrawRectangleRecC(r rl.Rectangle, col color.RGBA) {
	rl.DrawRectangleRec(rl.Rectangle{
		X:      r.X - r.Width/2,
		Y:      r.Y - r.Height/2,
		Width:  r.Width,
		Height: r.Height,
	}, col)
}

// Draw rectangle outline with extended parameters and anchor at its center
//
// If you're using this within normalized space, then pass `lineThick` as `thicknessValue * (2/winHeight)`
// to get the same thickness that you see in pixel space. 2 is the normalized height here.
func DrawRectangleLinesExC(r rl.Rectangle, lineThick float32, col color.RGBA) {
	rl.DrawRectangleLinesEx(rl.Rectangle{
		X:      r.X - r.Width/2,
		Y:      r.Y - r.Height/2,
		Width:  r.Width,
		Height: r.Height,
	}, lineThick, col)
}
