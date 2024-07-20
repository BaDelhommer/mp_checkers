package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Board struct {
	Pieces          [][]Piece
	WhitePiecesLeft int32
	RedPiecesLeft   int32
}

func (b *Board) CreateBoard() {
	for row := range ROWS {
		b.Pieces = append(b.Pieces, []Piece{})
		for col := range COLS {
			whitePiece := Piece{Color: rl.White, Row: row, Col: col, King: false}
			redPiece := Piece{Color: rl.Red, Row: row, Col: col, King: false}
			if col%2 == ((row + 1) % 2) {
				if row < 3 {
					b.Pieces[row] = append(b.Pieces[row], whitePiece)
					whitePiece.calcPosition()
					whitePiece.draw()
				} else if row > 4 {
					b.Pieces[row] = append(b.Pieces[row], redPiece)
					redPiece.calcPosition()
					redPiece.draw()
				} else {
					b.Pieces = append(b.Pieces, nil)
				}
			} else {
				b.Pieces = append(b.Pieces, nil)
			}
		}
	}
}

func drawSquares() {
	for row := range ROWS {
		for col := row % 2; col < ROWS; col += 2 {
			rl.DrawRectangle(row*SQUARE_SIZE, col*SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE, rl.Red)
		}
	}
}
