package main

import (
	"image/color"
	"log"
	"the-mountains/assets"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// TitleMenuScene shows a menu and waits for user input.
type TitleMenuScene struct {
	world *katsu2d.World
}

func NewTitleMenuScene(tm *katsu2d.TextureManager) *katsu2d.Scene {
	scene := katsu2d.NewScene()
	titleMenuScene := &TitleMenuScene{
		world: scene.World,
	}

	scene.AddSystem(titleMenuScene)
	scene.AddSystem(katsu2d.NewSpriteRenderSystem(tm))
	scene.AddSystem(katsu2d.NewTextRenderSystem())

	scene.OnEnter = titleMenuScene.OnEnter
	scene.OnExit = titleMenuScene.OnExit

	return scene
}

func (self *TitleMenuScene) OnEnter(e *katsu2d.Engine) {
	println("Entering TitleMenuScene...")

	// Add a sprite to the scene's world.
	sprite := self.world.CreateEntity()
	self.world.AddComponent(sprite, katsu2d.NewTransformComponent())
	s := katsu2d.NewSpriteComponent(0, 32, 32)
	s.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	self.world.AddComponent(sprite, s)
	tx, _ := self.world.GetComponent(sprite, katsu2d.CTTransform)
	t := tx.(*katsu2d.TransformComponent)
	t.SetPosition(ebimath.V2(100))

	// Add text for instructions.
	instructions := self.world.CreateEntity()
	self.world.AddComponent(instructions, katsu2d.NewTransformComponent())
	self.world.AddComponent(instructions,
		katsu2d.NewTextComponent(e.FontManager().Get(assets.AccidentalPresidencyFontID),
			"Press Enter to start!", 24, color.RGBA{R: 255, G: 255, B: 255, A: 255}).
			SetAlignment(katsu2d.TextAlignmentMiddleCenter))

	// Center the splash text.
	itx, _ := self.world.GetComponent(instructions, katsu2d.CTTransform)
	it := itx.(*katsu2d.TransformComponent)
	it.SetPosition(ebimath.V(float64(screenWidth)/2, float64(screenHeight)/2))
}

func (self *TitleMenuScene) OnExit(e *katsu2d.Engine) {
	println("Exiting TitleMenuScene...")
}

func (self *TitleMenuScene) Update(w *katsu2d.World, dt float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Println("User pressed Enter. This would start the game.")
		// We could switch to a "GameScene" here.
	}
}
