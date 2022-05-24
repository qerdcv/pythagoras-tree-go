package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	treeInitialHeight = 200.00
	branchWidth       = 20

	iterations = 4
)

var (
	branchColor = color.RGBA{G: 255, A: 255}
)

type tree struct {
}

func newTree() *tree {
	return &tree{}
}

func (t *tree) draw(screen *ebiten.Image) {
	initX := float64(screen.Bounds().Dx() / 2)
	initY := float64(screen.Bounds().Dy())

	// draw root
	ebitenutil.DrawLine(screen, initX, initY, initX, initY-treeInitialHeight, branchColor)

	t.drawBranch(screen, initX, initY, initX, initY-treeInitialHeight, treeInitialHeight, 0)
	// draw left side
	// t.drawLeftBranch(screen, initX, initY, initX, initY-treeInitialHeight, treeInitialHeight, 0)
	// t.drawRightBranch(screen, initX, initY, initX, initY-treeInitialHeight, treeInitialHeight, 0)
}

func (t *tree) drawBranch(screen *ebiten.Image, x1, y1, x2, y2, h, n float64) {
	if n >= iterations {
		// os.Exit(0)
		return
	}

	h /= 2

	x1 = x2
	y1 = y2

	n += 1

	dx := h * math.Cos(degToRad(45*n))
	dy := h * math.Sin(degToRad(45*n))

	ebitenutil.DrawLine(screen, x1, y1, x1-dx, y1-dy, branchColor)
	ebitenutil.DrawLine(screen, x1, y1, x1+dx, y1-dy, branchColor)

	t.drawBranch(screen, x1, y1, x1-dx, y1-dy, h, n)
	t.drawBranch(screen, x1, y1, x1+dx, y1-dy, h, n)
}

func (t *tree) drawLeftBranch(screen *ebiten.Image, x1, y1, x2, y2, h, n float64) {
	if n >= iterations {
		return
	}

	h /= 2

	x1 = x2
	y1 = y2

	n += 1

	dx := h * math.Sin(degToRad(45*n))
	dy := h * math.Cos(degToRad(45*n))

	x2 = x1 - dx
	y2 = y1 - dy

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, branchColor)

	// t.drawLeftBranch(screen, x1, y1, x2, y2, h, n)
	// t.drawRightBranch(screen, x1, y1, x2, y2, h, n)
}

func (t *tree) drawRightBranch(screen *ebiten.Image, x1, y1, x2, y2, h, n float64) {
	if n >= iterations {
		// os.Exit(0)
		return
	}

	h /= 2

	x1 = x2
	y1 = y2

	n += 1

	dx := h * math.Cos(degToRad(45*n))
	dy := h * math.Sin(degToRad(45*n))

	x2 = x1 + dx
	y2 = y1 - dy

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, branchColor)

	t.drawLeftBranch(screen, x1, y1, x2, y2, h, n)
	t.drawRightBranch(screen, x1, y1, x2, y2, h, n)
}

func (t *tree) update() {

}

func flipLine(x1, y1, x2, y2 float64, deg float32) {

}
