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

func NewGame() *Game {
	return &Game{
		Board:      Board{},
		Selected:   nil,
		Turn:       rl.Red,
		ValidMoves: [][]*Piece{},
	}
}
