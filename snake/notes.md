Line 25-29 in 6acc21caa6ef8663cfca6bcd244214acd93f6daf contains logic about moving the snake head. I found the frame time dependent logic interesting.  
Add this line after 27, to make snake's head wrap around

```go
snakeHeadPos.Y = float32(math.Mod(float64(snakeHeadPos.Y), GridWidth))
```
