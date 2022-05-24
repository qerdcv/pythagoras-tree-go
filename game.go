package main

import "github.com/hajimehoshi/ebiten/v2"

type game struct {
	width, height int
	*tree
}

func NewGame(width, height int) *game {
	return &game{width, height, newTree()}
}

func (g *game) Draw(screen *ebiten.Image) {
	g.tree.draw(screen)
}

func (g *game) Update() error {
	g.tree.update()
	return nil
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
