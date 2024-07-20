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
		rl.ClearBackground(rl.Black)
		drawSquares()

		rl.EndDrawing()
	}
}
