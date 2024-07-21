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

func (b *Board) traverseLeft(start, stop, step, left int32, color rl.Color, skipped []*Piece) [][]*Piece {
	moves := [][]*Piece{}
	last := []*Piece{}

	for i := start; i <= stop; i += step {
		if left < 0 {
			break
		}

		current := b.Pieces[i][left]
		if current.Empty {
			if len(skipped) > 0 && len(last) == 0 {
				break
			} else if len(skipped) > 0 {
				moves[i] = append(last, skipped...)
			} else {
				moves[i] = last
			}

			if len(last) > 0 {
				if step == -1 {
					current.Row = max(i-3, 0)
				} else {
					current.Row = min(i+3, ROWS)
				}
				skipped = append(skipped, last...)
				moves = append(moves, b.traverseLeft(i+step, current.Row, step, left-1, color, skipped)...)
				moves = append(moves, b.traverseRight(i+step, current.Row, step, left+1, color, skipped)...)
			}
			break
		} else if current.Color == color {
			break
		} else {
			last = append(last, current)
		}
		left -= 1
	}
	return moves
}

func (b *Board) traverseRight(start, stop, step, right int32, color rl.Color, skipped []*Piece) [][]*Piece {
	moves := [][]*Piece{}
	last := []*Piece{}

	for i := start; i <= stop; i += step {
		if right >= COLS {
			break
		}

		current := b.Pieces[i][right]
		if current.Empty {
			if len(skipped) > 0 && len(last) == 0 {
				break
			} else if len(skipped) > 0 {
				moves[i] = append(last, skipped...)
			} else {
				moves[i] = last
			}

			if len(last) > 0 {
				if step == -1 {
					current.Row = max(i-3, 0)
				} else {
					current.Row = min(i+3, ROWS)
				}
				skipped = append(skipped, last...)
				moves = append(moves, b.traverseLeft(i+step, current.Row, step, right-1, color, skipped)...)
				moves = append(moves, b.traverseLeft(i+step, current.Row, step, right+1, color, skipped)...)
			}
			break
		} else if current.Color == color {
			break
		} else {
			last = append(last, current)
		}
		right -= 1
	}
	return moves
}

func drawSquares() {
	for row := range ROWS {
		for col := row % 2; col < ROWS; col += 2 {
			rl.DrawRectangle(row*SQUARE_SIZE, col*SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE, rl.Red)
		}
	}
}
