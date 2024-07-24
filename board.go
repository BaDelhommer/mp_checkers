package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	Pieces          [ROWS][COLS]*Piece
	WhitePiecesLeft int32
	RedPiecesLeft   int32
}

func (b *Board) getPiece(row, col int32) *Piece {
	return b.Pieces[row][col]
}

func (b *Board) CreateBoard() {
	for row := range ROWS {
		for col := range COLS {
			whitePiece := Piece{Color: rl.White, Row: row, Col: col, King: false, Empty: false}
			redPiece := Piece{Color: rl.Red, Row: row, Col: col, King: false, Empty: false}
			blankPiece := Piece{Color: rl.Blank, Row: row, Col: col, King: false, Empty: true}
			if col%2 == ((row + 1) % 2) {
				if row < 3 {
					b.Pieces[row][col] = &whitePiece
					whitePiece.calcPosition()
				} else if row > 4 {
					b.Pieces[row][col] = &redPiece
					redPiece.calcPosition()
				} else {
					b.Pieces[row][col] = &blankPiece
					blankPiece.calcPosition()
				}
			} else {
				b.Pieces[row][col] = &blankPiece
				blankPiece.calcPosition()
			}
		}
	}
}

func (b *Board) move(piece *Piece, row int32, col int32) {
	b.Pieces[row][col].PosX = piece.PosX
	b.Pieces[row][col].PosY = piece.PosY
	piece.move(row, col)

	if row == ROWS-1 || row == 0 {
		piece.makeKing()
	}
}

func (b *Board) Remove(pieces map[[2]int32][]*Piece) {
	for _, ps := range pieces {
		for _, piece := range ps {
			if piece.Color == rl.White {
				b.WhitePiecesLeft -= 1
			} else {
				b.RedPiecesLeft -= 1
			}
		}
	}
}

func (b *Board) traverseLeft(start, stop, step, left int32, color rl.Color, skipped []*Piece) map[[2]int32][]*Piece {
	moves := make(map[[2]int32][]*Piece)
	var last []*Piece

	for r := start; r != stop; r += step {
		if left < 0 {
			break
		}
		current := b.Pieces[r][left]
		if current.Empty {
			if len(skipped) > 0 && len(last) == 0 {
				break
			} else if len(skipped) > 0 {
				moves[[2]int32{r, left}] = append(last, skipped...)
			} else {
				moves[[2]int32{r, left}] = last
			}
			if len(last) > 0 {
				var row int32
				if step == -1 {
					row = int32(math.Max(float64(r+3), 0))
				} else {
					row = int32(math.Min(float64(r+3), float64(ROWS)))
				}
				leftMoves := b.mergeMaps(moves, b.traverseLeft(r+step, row, step, left-1, color, last))
				rightMoves := b.mergeMaps(moves, b.traverseRight(r+step, row, step, left+1, color, last))
				moves = b.mergeMaps(moves, leftMoves)
				moves = b.mergeMaps(moves, rightMoves)
			}
			break
		} else if current.Color == color {
			break
		} else {
			last = []*Piece{current}
		}
		left -= 1
	}
	return moves
}

func (b *Board) traverseRight(start, stop, step, right int32, color rl.Color, skipped []*Piece) map[[2]int32][]*Piece {
	moves := make(map[[2]int32][]*Piece)
	var last []*Piece

	for r := start; r != stop; r += step {
		if right >= COLS {
			break
		}

		current := b.Pieces[r][right]
		if current.Empty {
			if len(skipped) > 0 && len(last) == 0 {
				break
			} else if len(skipped) > 0 {
				moves[[2]int32{r, right}] = append(last, skipped...)
			} else {
				moves[[2]int32{r, right}] = last
			}

			if len(last) > 0 {
				var row int32
				if step == -1 {
					row = int32(math.Max(float64(r-3), 0))
				} else {
					row = int32(math.Max(float64(r+3), float64(ROWS)))
				}
				leftMoves := b.mergeMaps(moves, b.traverseLeft(r+step, row, step, right-1, color, last))
				rightMoves := b.mergeMaps(moves, b.traverseRight(r+step, row, step, right+1, color, last))
				moves = b.mergeMaps(moves, leftMoves)
				moves = b.mergeMaps(moves, rightMoves)
			}
			break
		} else if current.Color == color {
			break
		} else {
			last = []*Piece{current}
		}
		right += 1
	}
	return moves
}

func (b *Board) GetValidMoves(piece *Piece) map[[2]int32][]*Piece {
	moves := make(map[[2]int32][]*Piece)
	left := piece.Col - 1
	right := piece.Col + 1
	row := piece.Row

	if piece.Color == rl.Red || piece.King {
		leftMoves := b.traverseLeft(row-1, int32(math.Max(float64(row-3), float64(-1))), int32(-1), left, piece.Color, []*Piece{})
		rightMoves := b.traverseRight(row-1, int32(math.Max(float64(row-3), float64(-1))), int32(-1), right, piece.Color, []*Piece{})
		moves = b.mergeMaps(moves, leftMoves)
		moves = b.mergeMaps(moves, rightMoves)
	}

	if piece.Color == rl.White || piece.King {
		leftMoves := b.traverseLeft(row+1, int32(min(float64(row+3), float64(ROWS))), 1, left, piece.Color, []*Piece{})
		rightMoves := b.traverseRight(row+1, int32(min(float64(row+3), float64(ROWS))), 1, right, piece.Color, []*Piece{})
		moves = b.mergeMaps(moves, leftMoves)
		moves = b.mergeMaps(moves, rightMoves)
	}
	return moves
}

func (b *Board) draw() {
	for row := range ROWS {
		for col := range COLS {
			piece := b.Pieces[row][col]
			if !piece.Empty {
				piece.draw()
			}
		}
	}
}

func (b *Board) mergeMaps(dest, src map[[2]int32][]*Piece) map[[2]int32][]*Piece {
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

func drawSquares() {
	for row := range ROWS {
		for col := row % 2; col < ROWS; col += 2 {
			rl.DrawRectangle(row*SQUARE_SIZE, col*SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE, rl.Red)
		}
	}
}
