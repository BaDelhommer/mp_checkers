package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Board struct {
	Pieces          [][]*Piece
	WhitePiecesLeft int32
	RedPiecesLeft   int32
}

func (b *Board) getPiece(row, col int32) *Piece {
	return b.Pieces[row][col]
}

func (b *Board) CreateBoard() {
	for row := range ROWS {
		b.Pieces = append(b.Pieces, []*Piece{})
		for col := range COLS {
			whitePiece := Piece{Color: rl.White, Row: row, Col: col, King: false, Empty: false}
			redPiece := Piece{Color: rl.Red, Row: row, Col: col, King: false, Empty: false}
			blankPiece := Piece{Color: rl.Blank, Row: row, Col: col, King: false, Empty: true}
			if col%2 == ((row + 1) % 2) {
				if row < 3 {
					b.Pieces[row] = append(b.Pieces[row], &whitePiece)
					whitePiece.calcPosition()
					whitePiece.draw()
				} else if row > 4 {
					b.Pieces[row] = append(b.Pieces[row], &redPiece)
					redPiece.calcPosition()
					redPiece.draw()
				} else {
					b.Pieces[row] = append(b.Pieces[row], &blankPiece)
				}
			} else {
				b.Pieces[row] = append(b.Pieces[row], &blankPiece)
			}
		}
	}
}

func (b *Board) move(piece *Piece, row int32, col int32) {
	b.Pieces[row] = nil
	piece.Row = row
	piece.Col = col
	piece.move(row, col)

	if row == ROWS-1 || row == 0 {
		piece.makeKing()
	}
}

func (b *Board) Remove(pieces []Piece) {
	for _, piece := range pieces {
		if !piece.Empty {
			if piece.Color == rl.Red {
				b.RedPiecesLeft -= 1
			} else {
				b.WhitePiecesLeft -= 1
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
