package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Board
	Selected   *Piece
	Turn       rl.Color
	ValidMoves [][]*Piece
}

func (g *Game) Select(row, col int32) bool {
	if g.Selected != nil {
		result := g.move(row, col)
		if !result {
			g.Selected = nil
			g.Select(row, col)
		}
	}

	piece := g.Board.getPiece(row, col)
	if !piece.Empty && piece.Color == g.Turn {
		g.Selected = piece
		g.ValidMoves = g.Board.GetValidMoves(piece)
		return true
	}
	return false
}

func (g *Game) changeTurn() {
	if g.Turn == rl.Red {
		g.Turn = rl.White
	} else {
		g.Turn = rl.Red
	}
}

func (g *Game) move(row, col int32) bool {
	piece := g.Board.getPiece(row, col)
	if g.Selected != nil && !piece.Empty && isMoveValid(g.ValidMoves, piece) {
		g.Board.move(g.Selected, row, col)
		skipped := []*Piece{}
		skipped = append(skipped, g.ValidMoves[row][col])
		if len(skipped) > 0 {
			g.Board.Remove(skipped)
		}
		g.changeTurn()
	} else {
		return false
	}

	return true
}

func (g *Game) showValidMoves(moves [][]*Piece) {
	for i, move := range moves {
		rl.DrawCircle(move[i].PosX, move[i].PosY, float32(SQUARE_SIZE/4), rl.Blue)
	}
}

func NewGame() *Game {
	return &Game{
		Board:      Board{},
		Selected:   nil,
		Turn:       rl.Red,
		ValidMoves: [][]*Piece{},
	}
}

func isMoveValid(pieces [][]*Piece, piece *Piece) bool {
	for i, p := range pieces {
		if p[i] == piece {
			return true
		}
	}
	return false
}
