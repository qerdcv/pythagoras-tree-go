package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	initHeight = 200
	nesting    = 10
	delta      = 0.67
	rotDeg     = 125
)

var (
	treeColor = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
)

type tree struct {
	rx, ry float64
}

func newTree(screenWidht, screenHeight int) *tree {
	return &tree{
		float64(screenWidht) / 2,
		float64(screenHeight),
	}
}

func (t *tree) generateTree(screen *ebiten.Image, n, sx, sy, h float64) {
	t.generateLeftBranch(screen, n, -rotDeg, sx, sy, h)
	t.generateRightBranch(screen, n, rotDeg, sx, sy, h)
}

func (t *tree) generateLeftBranch(screen *ebiten.Image, n, deg, sx, sy, h float64) {
	if n >= nesting {
		return
	}

	lb := line{sx, sy, sx, sy - h, h}
	lb.rotate(deg)
	lb.draw(screen)

	h *= delta

	n++
	t.generateLeftBranch(screen, n, deg-rotDeg, lb.x2, lb.y2, h)
	t.generateRightBranch(screen, n, deg+rotDeg, lb.x2, lb.y2, h)
}

func (t *tree) generateRightBranch(screen *ebiten.Image, n, deg, sx, sy, h float64) {
	if n >= nesting {
		return
	}

	rb := line{sx, sy, sx, sy - h, h}
	rb.rotate(deg)
	rb.draw(screen)

	h *= delta

	n++
	t.generateRightBranch(screen, n, deg+rotDeg, rb.x2, rb.y2, h)
	t.generateLeftBranch(screen, n, deg-rotDeg, rb.x2, rb.y2, h)
}

func (t *tree) Draw(screen *ebiten.Image) {
	root := line{t.rx, t.ry, t.rx, t.ry - initHeight, initHeight}
	root.draw(screen)
	t.generateTree(screen, 1, root.x2, root.y2, initHeight/2)
}
