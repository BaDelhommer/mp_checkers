package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	*Board
	Selected   *Piece
	Turn       rl.Color
	ValidMoves map[[2]int32][]*Piece
}

func (g *Game) Select(row, col int32) bool {
	if g.Selected != nil {
		result := g.move(row, col)
		if !result {
			g.Selected.Empty = true
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
	if !g.Selected.Empty && !piece.Empty && isMoveValid(g.ValidMoves, piece) {
		fmt.Println("First condition")
		g.Board.move(g.Selected, row, col)
		var skipped map[[2]int32][]*Piece
		skipped = g.mergeMaps(skipped, g.ValidMoves)
		if len(skipped) > 0 {
			fmt.Println("Second condition")
			g.Board.Remove(skipped)
		}
		g.changeTurn()
	} else {
		return false
	}

	return true
}

func (g *Game) showValidMoves(moves map[[2]int32][]*Piece) {
	for c := range moves {
		rl.DrawCircle(c[1]*SQUARE_SIZE+(SQUARE_SIZE/2), c[0]*SQUARE_SIZE+(SQUARE_SIZE/2), float32(SQUARE_SIZE/4), rl.Blue)
	}
}

func NewGame() *Game {
	return &Game{
		Board:      &Board{},
		Selected:   &Piece{},
		Turn:       rl.Red,
		ValidMoves: map[[2]int32][]*Piece{},
	}
}

func isMoveValid(pieces map[[2]int32][]*Piece, piece *Piece) bool {
	for _, p := range pieces {
		for _, i := range p {
			if i == piece {
				return true
			}
		}
	}
	return false
}
