package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	leftRot  = -1
	rightRot = 1

	maxNesting = 5
)

var (
	treeColor = color.RGBA{
		R: 9,
		G: 255,
		B: 9,
		A: 255,
	}

	isCalculated = false
)

type tree struct {
}

func newTree() *tree {
	return &tree{}
}

func (t *tree) update() {
	if isCalculated {
		return
	}

	isCalculated = true

}

func (t *tree) draw(screen *ebiten.Image) {
	drawTree(screen, float64(screenWidth/2), screenHeight, screenHeight/2, 0, leftRot)
}

func drawTree(screen *ebiten.Image, x1, y1, h, n, rot float64) {
	if n == maxNesting {
		return
	}

	h /= 2
	x2 := h + h*math.Cos(degToRad(45)*n*rot)
	y2 := h + h*math.Sin(degToRad(45)*n*rot)

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, treeColor)

	n += 1

	drawTree(screen, x2, y2, h, n, rot)
}

func degToRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}
