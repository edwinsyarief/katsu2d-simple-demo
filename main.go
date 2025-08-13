package main

import (
	"fmt"
	"image/color"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth  = 640
	screenHeight = 480
)

// --- ENGINE-LEVEL SYSTEMS (DRAWN BEFORE AND AFTER THE SCENE) ---

// BackgroundSystem is a DrawSystem for the engine, rendering a simple background color.
type BackgroundSystem struct{}

func (self *BackgroundSystem) Draw(e *katsu2d.Engine, renderer *katsu2d.BatchRenderer) {
	// This system draws at the very bottom of the render stack.
	renderer.GetScreen().Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
}

// FPSSystem is a DrawSystem that draws the FPS counter as a global overlay.
type FPSSystem struct{}

func (self *FPSSystem) Draw(e *katsu2d.Engine, renderer *katsu2d.BatchRenderer) {
	// This system draws at the very top of the render stack.
	text := katsu2d.NewText(
		e.FontManager().Get(DefaultFontID),
		fmt.Sprintf("FPS %.2f", ebiten.ActualFPS()),
		30,
		color.RGBA{R: 255, G: 255, B: 255, A: 255})
	t := ebimath.T()
	t.SetPosition(ebimath.V(10, 10))
	renderer.DrawText(text, t)
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
		katsu2d.WithOverlayDrawSystem(&FPSSystem{}),
	)

	game.InitFS(FS)
	loadAssets(game)
	game.SetTimeScale(1)

	// Create and add the scenes to the scene manager.
	game.AddScene("intro", NewIntroScene())
	game.AddScene("titleMenu", NewTitleMenuScene())

	// Start the game by switching to the first scene.
	game.SwitchScene("intro")

	return game
}

func loadAssets(e *katsu2d.Engine) {
	// textures
	EbitengineLogoTextureID = e.TextureManager().LoadEmbedded("assets/images/ebitengine_logo.png")

	// fonts
	DefaultFontID = e.FontManager().LoadEmbedded("assets/fonts/default.ttf")
}

func runGame(game *katsu2d.Engine) {
	if err := game.Run(); err != nil {
		panic(err)
	}
}
