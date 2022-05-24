package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 680
	screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)

	ebiten.SetWindowTitle("Pythagoras Tree")
	if err := ebiten.RunGame(NewGame(screenWidth, screenHeight)); err != nil {
		log.Fatal(err)
	}
}
