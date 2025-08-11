package main

import (
	"image/color"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Initialize engine
	game := katsu2d.NewEngine(
		katsu2d.WithWindowSize(640, 480),
		katsu2d.WithWindowTitle("Katsu2D Simple Demonstration"),
		katsu2d.WithWindowResizeMode(ebiten.WindowResizingModeEnabled),
		katsu2d.WithClearScreenEachFrame(false),
	)

	// Register built-in system
	game.AddSystem(
		katsu2d.NewRenderSystem(game.TextureManager()),
	)

	// Create a new entity
	e := game.World().NewEntity()

	// Create a transform component
	t := katsu2d.NewTransform()
	t.SetPosition(ebimath.V(100, 100))

	// Create a sprite component
	s := katsu2d.NewSprite(0, 50, 50)
	s.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// Add components into entity
	game.World().AddComponent(e, t)
	game.World().AddComponent(e, s)

	if err := game.Run(); err != nil {
		panic(err)
	}
}
