package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Hello from abhi")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var font_size int32 = 32
	font := rl.LoadFontEx("./FiraSans-Regular.otf", font_size, nil)

	rl.SetExitKey(rl.KeyQ)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		dimen := rl.MeasureTextEx(font, "Welcome! Press q to exit", float32(font_size), 0)

		pos := rl.Vector2{
			X: 800/2 - dimen.X/2,
			Y: 450/2 - dimen.Y/2,
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			rl.CloseWindow()
		}
		rl.DrawTextEx(font, "Welcome! Press q to exit", pos, float32(font_size), 0, rl.RayWhite)
		rl.EndDrawing()
	}

}
