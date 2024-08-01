## Install dependencies

```sh
sudo dnf install SDL2-devel
cd project-root
go mod tidy # or `go get .`
```

Using SDL over GLFW for Raylib because in Wayland it:

- lets me position the window to a monitor
- doesn't has wonky rendering
- doesn't has scaling/positioning issues

## Running a program

```sh
cd into-a-study
go run -tags sdl .
```

## Watching for changes

```sh
watchexec -qnrc -- go run -tags sdl .
```

IF you want to watch some other directory too, then use:

```sh
watchexec -qnrc -w ../util -w . -- go run -tags sdl .
```

## Loading an external font

```go
// https://github.com/adrg/sysfont
fontFinder := sysfont.NewFinder(&sysfont.FinderOpts{
	Extensions: []string{".ttf"},
})
matchedFont := fontFinder.Match("Segoe UI")
if matchedFont == nil {
	log.Fatal("system font not found")
}
font := rl.LoadFontEx(matchedFont.Filename, 18, []rune(nil)) // Pass nil to load the default character set
defer rl.UnloadFont(font)

// Following code is not needed if you aren't using FilterTrilinear. On 2D drawing FilterTrilinear isn't noticeable because it looks like FilterBiLinear
// rl.GenTextureMipmaps(&font.Texture)
// rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)
```

## Centering text:

```go
if loading {
	text := "Retrieving brightness..."
	var spacing float32 = 1
	textSize := rl.MeasureTextEx(font, text, fontSize, spacing)
	centeredPosition := rl.Vector2{
		X: float32(WinWidth)/2 - textSize.X/2,
		Y: float32(WinHeight)/2 - textSize.Y/2,
	}
	rl.DrawTextEx(font, text, centeredPosition, fontSize, spacing, rl.LightGray)

	rl.EndDrawing()
	continue
}
```

You can add some negative value to Y if think the font not optically centered. I'm unsure if Raylib supports [leading trim](https://medium.com/microsoft-design/leading-trim-the-future-of-digital-typesetting-d082d84b202) and if not, how easy it would be to support it.

## Variable fonts

See https://claude.ai/chat/ba51903f-ab2e-445e-b26c-b8012009d123

## Ensuring only a single window of the instance is ever launched

- [In Windows](https://github.com/ajitid/stellate/blob/main/single-instance_windows.go)
- [In macOS and Linux](https://claude.ai/chat/38e56e68-e64a-4a1b-8272-7ac1a5e7ba82)

## Do movements/animation using timer

```go
rl.SetConfigFlags(rl.FlagVsyncHint)
rl.InitWindow(WinSize, WinSize, "snake")
defer rl.CloseWindow()

for !rl.WindowShouldClose() {
	tickTimer -= rl.GetFrameTime()
	if tickTimer <= 0 {
		snakeHeadPos.Y += 1 // your movement changes go here
		tickTimer = TickRate + tickTimer
	}

	rl.BeginDrawing()
	// ...
	rl.EndDrawing()
}
```

See a concrete implementation in commit 5d726bff5cc6405e6aed574a4e90af148f1a5979

## Getting focus back on editor

Because SDL doesn't support `rl.FlagWindowUnfocused` at the moment, as soon as the renderer window is open it grabs the focus away from editor. Which is probably not what you want if you are tweaking some values to see what happens.

To fix this in Wayland Gnome, install [this extension](https://extensions.gnome.org/extension/5021/activate-window-by-title/).

If you install it directly, you may need to hit `alt-f2`, type `r` and hit enter.
Otherwise you can install it using [Extension Manager](https://github.com/mjakeman/extension-manager) app.

Call it as a goroutine in your main(), preferably after `rl.SetWindowMonitor()` if you've used it:

```go
go util.FocusEditor()
```

Further reading:

- https://unix.stackexchange.com/a/700116

## Libs

- Alternative to rlgl/OpenGL to make shaders in Go https://github.com/gopxl/glhf
- Sound https://github.com/gopxl/beep
- Retry github.com/avast/retry-go/v4
- Global hotkey github.com/robotn/gohook. For special keys, [see this](https://github.com/ajitid/stellate/blob/06989b0de27999ff514d87a959bcf8a147904693/main.go#L132-L152)
- Keys emulation github.com/micmonay/keybd_event
- Tween animation https://github.com/tanema/gween
- Spring https://github.com/charmbracelet/harmonica
- Physics engine
  - Basic
    - https://github.com/divVerent/awesome-ebiten?tab=readme-ov-file#physics
    - https://github.com/rudransh61/Physix-go
  - Rope
    - https://x.com/hemarkable/status/1817105134243184999
    - https://github.com/ByteArena/box2d/blob/master/DynamicsB2Rope.go
  - Nature
    - https://pkg.go.dev/github.com/g3n/engine/experimental/physics
    - https://commerce.nearform.com/open-source/renature

## Other resources

- [Notes](https://github.com/ajitid/stellate/blob/main/notes-raylib.md?plain=1#L44)
