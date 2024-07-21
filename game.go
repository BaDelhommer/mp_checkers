package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Board
	Selected   *Piece
	Turn       rl.Color
	ValidMoves [][]*Piece
}
