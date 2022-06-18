package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	*tree
	width, height int
}

func NewGame(width, height int) *Game {
	return &Game{
		tree:   newTree(),
		width:  width,
		height: height,
	}
}

func (g Game) Update() error {
	g.tree.update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.tree.draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
