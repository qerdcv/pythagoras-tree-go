package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*tree
	*rotatingLine
	width, height float64
}

func NewGame(width, height float64) *Game {
	return &Game{
		tree:         newTree(),
		rotatingLine: newRotatingLine(width/2, height/2),
		width:        width,
		height:       height,
	}
}

func (g Game) Update() error {
	g.tree.update()
	// g.rotatingLine.update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {

	// g.rotatingLine.d
	// raw(screen)
	g.tree.draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
