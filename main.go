package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Hello world")

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "checkers game")
	defer rl.CloseWindow()
	game_board := Board{
		Pieces:          [][]*Piece{},
		WhitePiecesLeft: NUM_PIECES,
		RedPiecesLeft:   NUM_PIECES,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawSquares()
		game_board.CreateBoard()

		rl.EndDrawing()
	}
}
