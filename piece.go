package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Piece struct {
	Color rl.Color
	Col   int32
	Row   int32
	King  bool
	PosX  int32
	PosY  int32
	Empty bool
}

func (p *Piece) makeKing() {
	p.King = true
}

func (p *Piece) calcPosition() {
	p.PosX = SQUARE_SIZE*p.Col + SQUARE_SIZE/2
	p.PosY = SQUARE_SIZE*p.Row + SQUARE_SIZE/2
}

func (p *Piece) move(row, col int32) {
	p.Row = row
	p.Col = col
	p.calcPosition()
}

func (p *Piece) draw() {
	radius := SQUARE_SIZE/2 - PADDING
	rl.DrawCircle(p.PosX, p.PosY, float32(radius+OUTLINE), rl.Gray)
	rl.DrawCircle(p.PosX, p.PosY, float32(radius), p.Color)

	if p.King {
		fmt.Println("KING ME!!!")
	}
}
