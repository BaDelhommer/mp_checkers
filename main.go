package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Hello world")

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "checkers game")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawText("Hello WOrld", WINDOW_WIDTH/2, WINDOW_HEIGHT/2, 16, rl.Black)
		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
