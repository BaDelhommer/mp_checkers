package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "checkers game")
	defer rl.CloseWindow()
	game := NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawSquares()
		game.Board.CreateBoard()
		game.Board.draw()
		if rl.IsMouseButtonDown(0) {
			x := rl.GetMouseX()
			y := rl.GetMouseY()
			row, col := getPieceRowCol(x, y)
			if game.Turn == rl.Red {
				game.Select(row, col)
				game.showValidMoves(game.ValidMoves)
				// fmt.Println(game.ValidMoves)
			}
			if game.Turn == rl.White {
				game.Select(row, col)
				game.showValidMoves(game.ValidMoves)
			}

		}

		rl.EndDrawing()
	}
}

func getPieceRowCol(x, y int32) (int32, int32) {
	row := y / SQUARE_SIZE
	col := x / SQUARE_SIZE

	return row, col
}
