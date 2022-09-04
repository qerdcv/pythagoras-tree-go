package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type line struct {
	x1, y1, x2, y2, length float64
}

func (l *line) rotate(deg float64) {
	s, c := math.Sincos(degToRad(deg))

	l.x2 = l.x1 - (l.length * c)
	l.y2 = l.y1 - (l.length * s)
}

func (l *line) draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, l.x1, l.y1, l.x2, l.y2, treeColor)
}
