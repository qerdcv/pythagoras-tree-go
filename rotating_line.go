package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const lineHeight = 200

var (
	lineColor = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
)

type rotatingLine struct {
	x1, y1, x2, y2, rot float64
}

func newRotatingLine(x1, y1 float64) *rotatingLine {
	return &rotatingLine{
		x1:  x1,
		y1:  y1,
		x2:  x1,
		y2:  y1 - lineHeight,
		rot: -90,
	}
}

func (rl *rotatingLine) update() {
	rl.rot++
	//
	// rl.x2 *= math.Cos(degToRad(90))
	// rl.y2 *= math.Sin(degToRad(90))

	dx := lineHeight * math.Cos(degToRad(rl.rot))
	dy := lineHeight * math.Sin(degToRad(rl.rot))

	rl.x2 = rl.x1 + dx
	rl.y2 = rl.y1 + dy
}

func (rl rotatingLine) draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, rl.x1, rl.y1, rl.x2, rl.y2, lineColor)
}
