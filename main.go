package main

import (
	"fmt"
	"image/color"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/edwinsyarief/katsu2d/utils"
	"github.com/edwinsyarief/lazyecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	screenWidth  = 800
	screenHeight = 600
)

// --- ENGINE-LEVEL SYSTEMS (DRAWN BEFORE AND AFTER THE SCENE) ---

// BackgroundSystem is a DrawSystem for the engine, rendering a simple background color.
type BackgroundSystem struct{}

func (self *BackgroundSystem) Draw(_ *lazyecs.World, renderer *katsu2d.BatchRenderer) {
	// This system draws at the very bottom of the render stack.
	renderer.GetScreen().Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
}

// FPSSystem is a DrawSystem that draws the FPS counter as a global overlay.
type FPSSystem struct {
	engine *katsu2d.Engine
}

func NewFPSSystem(engine *katsu2d.Engine) *FPSSystem {
	return &FPSSystem{
		engine: engine,
	}
}

func (self *FPSSystem) Draw(world *lazyecs.World, renderer *katsu2d.BatchRenderer) {
	// This system draws at the very top of the render stack.
	txt := katsu2d.NewDefaultTextComponent(
		fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()),
		30,
		color.RGBA{R: 255, G: 255, B: 255, A: 255})
	t := ebimath.T()
	t.SetPosition(ebimath.V(10, 10))

	op := &text.DrawOptions{}
	op.GeoM = t.Matrix()
	op.ColorScale = utils.RGBAToColorScale(txt.Color)
	text.Draw(renderer.GetScreen(), txt.Caption, txt.FontFace, op)
}

func main() {
	game := setupGame()
	runGame(game)
}

func setupGame() *katsu2d.Engine {
	game := katsu2d.NewEngine(
		katsu2d.WithWindowSize(screenWidth, screenHeight),
		katsu2d.WithWindowTitle("Katsu2D Simple Demo"),
		katsu2d.WithWindowResizeMode(ebiten.WindowResizingModeEnabled),
		katsu2d.WithFullScreen(false),
		katsu2d.WithVsyncEnabled(true),
		katsu2d.WithCursorMode(ebiten.CursorModeVisible),
		katsu2d.WithClearScreenEachFrame(false),
		katsu2d.WithBackgroundDrawSystem(&BackgroundSystem{}),
	)

	game.AddOverlayDrawSystem(NewFPSSystem(game))

	game.InitFS(FS)
	loadAssets(game)
	game.SetTimeScale(1)

	// Create and add the scenes to the scene manager.
	game.AddScene("intro", NewIntroScene())
	game.AddScene("titleMenu", NewTitleMenuScene(game.TextureManager()))

	// Start the game by switching to the first scene.
	game.SwitchScene("intro")

	return game
}

func loadAssets(e *katsu2d.Engine) {
	// textures
	EbitengineLogoTextureID = e.TextureManager().LoadEmbedded("assets/images/ebitengine_logo.png")
}

func runGame(game *katsu2d.Engine) {
	if err := game.Run(); err != nil {
		panic(err)
	}
}
