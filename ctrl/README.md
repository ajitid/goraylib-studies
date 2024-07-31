## How to use

### Setting up

Firstly, put `ctrl.IsDrawing` at the right places:

```go
for !rl.WindowShouldClose() {
  update()

  rl.BeginDrawing()
  ctrl.IsDrawing = true

  draw()

  ctrl.IsDrawing = false
  rl.EndDrawing()
```

Then right at the start of your main, run `ctrl.RunServer()` in a goroutine:

```go
func main() {
  go ctrl.RunServer()

  // ...
}
```

### Variable outside loop

If the variable is defined outside the for loop, then you can use it like so:

```go
var length float32 = 33.5
ctrl.SetFloat("length of red line", &length)
```

### Variable inside drawing loop

If the variable is defined inside the for loop however, it'd mean that every iteration will destory the variable. Hence keeping a pointer won't make sense. In this scenario, use the methods suffixed by `V`:

```go
points := ctrl.SetIntV("points", 32.34)
```

If you are defining a control this way, you won't be able to control the variable's value from Go.

PS: see commit 8ab3496bd36ee0dd3e2dba48d17b54cdab690616. We can expose `GetIntPtr()` and `informUI()` to the user and let and they can call these fns to manipulate and then send the value to the UI. (We can't use `SetInt()` as it is aware of the drawing loop and thus will refuse to update the value.) While it is technically possible to modify a value that was set using `SetIntV()`, there is no practical use for it.

### Variable inside update loop

Irrespective of whether you have created a variable within `update()` loop or outside it, you may want to visualise the change in its value over time. You should use a usual setter in this case:

```go
var disturbance float32 = 33.5
ctrl.SetFloat("disturbance", &disturbance)
```

### Wrong usages

You can't create 2 controls with the same name:

```go
var lengthRedLine float32 = 33.5
ctrl.SetFloat("length", &lengthRedLine)
var lengthPinkLine float32 = 33.5
ctrl.SetFloat("length", &lengthPinkLine) // ❌ Wrong usage. Control to lengthRedLine would be lost as lengthPinkLine will override it
```
