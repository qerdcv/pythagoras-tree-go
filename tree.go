package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	maxNesting = 5

	initSize = 200
	initRot  = -135
)

var (
	leftBranchColor = color.RGBA{
		R: 9,
		G: 255,
		B: 9,
		A: 255,
	}
	rightBranchColor = color.RGBA{
		R: 255,
		G: 9,
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
	x1 := float64(screenWidth / 2)
	y1 := float64(screenHeight)

	x2 := x1
	y2 := y1 - initSize

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, leftBranchColor)

	drawLeftBranch(screen, x2, y2, screenHeight/2, 0, initRot)
	// drawRightBranch(screen, x2, y2, screenHeight/2, 0, initRot)
}

func drawLeftBranch(screen *ebiten.Image, x1, y1, h, n, rotDeg float64) {
	if n == maxNesting {
		return
	}

	h /= 2

	dx := h * math.Cos(degToRad(rotDeg))
	dy := h * math.Sin(degToRad(rotDeg))

	x2 := x1 + dx
	y2 := y1 + dy

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, leftBranchColor)

	drawLeftBranch(screen, x2, y2, h, n+1, rotDeg-45)
	drawRightBranch(screen, x2, y2, h, n, rotDeg)
}

func drawRightBranch(screen *ebiten.Image, x1, y1, h, n, rotDeg float64) {
	if n == maxNesting {
		return
	}

	h /= 2

	dx := h * math.Cos(degToRad(rotDeg))
	dy := h * math.Sin(degToRad(rotDeg))

	x2 := x1 - dx
	y2 := y1 + dy

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, rightBranchColor)

	drawLeftBranch(screen, x2, y2, h, n+1, rotDeg+45)
	drawRightBranch(screen, x2, y2, h, n+1, rotDeg-45)
}

func degToRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}
