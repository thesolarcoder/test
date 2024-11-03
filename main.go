package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Hello from abhi")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		len := rl.MeasureText("Welcome!", 50)
		rl.DrawText("Welcome!", 800/2-len/2, 450/2-50/2, 50, rl.RayWhite)
		rl.EndDrawing()
	}

}
