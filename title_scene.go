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

func NewTitleMenuScene() *katsu2d.Scene {
	scene := katsu2d.NewScene()
	titleMenuScene := &TitleMenuScene{
		world: scene.World,
	}

	scene.AddSystem(titleMenuScene)
	scene.AddSystem(&SpriteDrawSystem{})
	scene.AddSystem(&TextDrawSystem{})

	scene.OnEnter = titleMenuScene.OnEnter
	scene.OnExit = titleMenuScene.OnExit

	return scene
}

func (self *TitleMenuScene) OnEnter(e *katsu2d.Engine) {
	println("Entering TitleMenuScene...")

	// Add a sprite to the scene's world.
	sprite := self.world.CreateEntity()
	self.world.AddComponent(sprite, katsu2d.NewTransform())
	s := katsu2d.NewSprite(0, 32, 32)
	s.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	self.world.AddComponent(sprite, s)
	tx, _ := self.world.GetComponent(sprite, katsu2d.CTTransform)
	t := tx.(*katsu2d.Transform)
	t.SetPosition(ebimath.V2(100))

	// Add text for instructions.
	instructions := self.world.CreateEntity()
	self.world.AddComponent(instructions, katsu2d.NewTransform())
	self.world.AddComponent(instructions,
		katsu2d.NewText(e.FontManager().Get(assets.AccidentalPresidencyFontID),
			"Press Enter to start!", 24, color.RGBA{R: 255, G: 255, B: 255, A: 255}).
			SetAlignment(katsu2d.TextAlignmentMiddleCenter))

	// Center the splash text.
	itx, _ := self.world.GetComponent(instructions, katsu2d.CTTransform)
	it := itx.(*katsu2d.Transform)
	it.SetPosition(ebimath.V(float64(screenWidth)/2, float64(screenHeight)/2))
}

func (self *TitleMenuScene) OnExit(e *katsu2d.Engine) {
	println("Exiting TitleMenuScene...")
}

func (self *TitleMenuScene) Update(e *katsu2d.Engine, dt float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Println("User pressed Enter. This would start the game.")
		// We could switch to a "GameScene" here.
	}
}

// SpriteDrawSystem is a DrawSystem for rendering sprite entities.
type SpriteDrawSystem struct{}

func (self *SpriteDrawSystem) Draw(e *katsu2d.Engine, renderer *katsu2d.BatchRenderer) {
	for _, entity := range e.SceneManager().Query(katsu2d.CTSprite, katsu2d.CTTransform) {
		tx, _ := e.SceneManager().GetComponent(entity, katsu2d.CTTransform)
		t := tx.(*katsu2d.Transform)
		sprite, _ := e.SceneManager().GetComponent(entity, katsu2d.CTSprite)
		s := sprite.(*katsu2d.Sprite)

		img := e.TextureManager().Get(s.TextureID)
		renderer.DrawQuad(
			t.Position(), t.Scale(), t.Offset(), t.Origin(), t.Rotation(),
			img, s.Color)
	}
}
