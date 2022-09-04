package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidht  = 640
	screenHeight = 480
)

type game struct {
	width  int
	height int
	title  string

	tree *tree
}

func newGame() *game {
	return &game{
		width:  screenWidht,
		height: screenHeight,
		title:  "Pythogoras tree",

		tree: newTree(screenWidht, screenHeight),
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.tree.Draw(screen)
}

func (g *game) Layout(outsizeWidth, outsideHeight int) (int, int) {
	return outsizeWidth, outsideHeight
}

func (g *game) Run() error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	return ebiten.RunGame(g)
}
