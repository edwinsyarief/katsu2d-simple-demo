package main

import (
	"image/color"
	"log"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/edwinsyarief/lazyecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// TitleMenuScene shows a menu and waits for user input.
type TitleMenuScene struct {
	world *lazyecs.World
}

func NewTitleMenuScene(tm *katsu2d.TextureManager) *katsu2d.Scene {
	scene := katsu2d.NewScene()
	titleMenuScene := &TitleMenuScene{
		world: scene.World,
	}

	scene.AddSystem(titleMenuScene)
	ls := katsu2d.NewLayerSystem(640, 480,
		katsu2d.AddSystem(katsu2d.NewSpriteRenderSystem(tm)),
		katsu2d.AddSystem(katsu2d.NewTextRenderSystem()),
	)
	scene.AddSystem(ls)

	scene.OnEnter = titleMenuScene.OnEnter
	scene.OnExit = titleMenuScene.OnExit

	return scene
}

func (self *TitleMenuScene) OnEnter(e *katsu2d.Engine) {
	println("Entering TitleMenuScene...")

	// Add a sprite to the scene's world.
	sprite := self.world.CreateEntity()

	t := katsu2d.TransformComponent{}
	t.Init()
	t.SetPosition(ebimath.V2(100))
	t.SetOffset(ebimath.V2(25))
	lazyecs.SetComponent(self.world, sprite, t)

	img := e.TextureManager().Get(0)
	s := katsu2d.SpriteComponent{}
	s.Init(0, img.Bounds())
	s.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	s.DstW = 50
	s.DstH = 50
	lazyecs.SetComponent(self.world, sprite, s)

	// Add text for instructions.
	instructions := self.world.CreateEntity()

	it := katsu2d.TransformComponent{}
	it.Init()
	it.SetPosition(ebimath.V(float64(640)/2, float64(480)/2))
	lazyecs.SetComponent(self.world, instructions, it)

	lazyecs.SetComponent(self.world, instructions, *katsu2d.NewDefaultTextComponent(
		"Press Enter to start!", 25, color.RGBA{R: 255, G: 255, B: 255, A: 255}).
		SetAlignment(katsu2d.TextAlignmentMiddleCenter))

}

func (self *TitleMenuScene) OnExit(e *katsu2d.Engine) {
	println("Exiting TitleMenuScene...")
}

func (self *TitleMenuScene) Update(w *lazyecs.World, dt float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Println("User pressed Enter. This would start the game.")
		// We could switch to a "GameScene" here.
	}

	query := w.Query(katsu2d.CTSprite)
	for query.Next() {
		for _, entity := range query.Entities() {
			t, _ := lazyecs.GetComponent[katsu2d.TransformComponent](self.world, entity)
			t.Rotate(ebimath.ToRadians(0.5))
		}
	}
}
