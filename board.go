package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawSquares() {
	for row := range ROWS {
		for col := row % 2; col < ROWS; col += 2 {
			rl.DrawRectangle(row*SQUARE_SIZE, col*SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE, rl.Red)
		}
	}
}
