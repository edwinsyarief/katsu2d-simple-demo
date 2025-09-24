package main

import (
	"image/color"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
	"github.com/edwinsyarief/lazyecs"
)

// IntroScene shows a splash screen and automatically transitions to the next scene.
type IntroScene struct {
	world      *lazyecs.World
	introTimer float64
	engine     *katsu2d.Engine
}

func NewIntroScene() *katsu2d.Scene {
	scene := katsu2d.NewScene()
	introScene := &IntroScene{
		world:      scene.World,
		introTimer: 0,
	}

	// Add the systems specific to this scene.
	scene.AddSystem(introScene)
	ls := katsu2d.NewLayerSystem(screenWidth, screenHeight,
		katsu2d.AddSystem(katsu2d.NewTextRenderSystem()),
	)
	scene.AddSystem(ls)

	// Add the lifecycle hooks.
	scene.OnEnter = introScene.OnEnter

	return scene
}

// OnEnter is called when the scene becomes active.
func (self *IntroScene) OnEnter(e *katsu2d.Engine) {
	self.engine = e
	println("Entering IntroScene...")

	// Create a splash screen text entity.
	splashText := self.world.CreateEntity()

	lazyecs.AddComponent[katsu2d.TransformComponent](self.world, splashText)

	t, _ := lazyecs.GetComponent[katsu2d.TransformComponent](self.world, splashText)
	t.Init()
	t.SetPosition(ebimath.V(float64(screenWidth)/2, float64(screenHeight)/2))

	lazyecs.SetComponent(self.world, splashText, *katsu2d.NewDefaultTextComponent(
		"Loading...", 35, color.RGBA{R: 255, G: 255, B: 255, A: 255}).
		SetAlignment(katsu2d.TextAlignmentMiddleCenter))

}

// Update is an UpdateSystem for the IntroScene.
func (self *IntroScene) Update(w *lazyecs.World, dt float64) {
	// Count down a timer.
	self.introTimer += dt
	if self.introTimer >= 3.0 { // Wait for 3 seconds.
		println("Transitioning to TitleMenuScene...")
		self.engine.SwitchScene("titleMenu")
	}
}
