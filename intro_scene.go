package main

import (
	"image/color"
	"the-mountains/assets"

	ebimath "github.com/edwinsyarief/ebi-math"
	"github.com/edwinsyarief/katsu2d"
)

// IntroScene shows a splash screen and automatically transitions to the next scene.
type IntroScene struct {
	world      *katsu2d.World
	introTimer float64
}

func NewIntroScene() *katsu2d.Scene {
	scene := katsu2d.NewScene()
	introScene := &IntroScene{
		world:      scene.World,
		introTimer: 0,
	}

	// Add the systems specific to this scene.
	scene.AddSystem(introScene)
	scene.AddSystem(&TextDrawSystem{})

	// Add the lifecycle hooks.
	scene.OnEnter = introScene.OnEnter

	return scene
}

// OnEnter is called when the scene becomes active.
func (self *IntroScene) OnEnter(e *katsu2d.Engine) {
	println("Entering IntroScene...")

	// Create a splash screen text entity.
	splashText := self.world.CreateEntity()
	self.world.AddComponent(splashText, katsu2d.NewTransform())
	self.world.AddComponent(splashText,
		katsu2d.NewText(e.FontManager().Get(assets.AccidentalPresidencyFontID),
			"Loading...", 48, color.RGBA{R: 255, G: 255, B: 255, A: 255}).
			SetAlignment(katsu2d.TextAlignmentMiddleCenter))

	// Center the splash text.
	tx, _ := self.world.GetComponent(splashText, katsu2d.CTTransform)
	t := tx.(*katsu2d.Transform)
	t.SetPosition(ebimath.V(float64(screenWidth)/2, float64(screenHeight)/2))
}

// Update is an UpdateSystem for the IntroScene.
func (self *IntroScene) Update(e *katsu2d.Engine, dt float64) {
	// Count down a timer.
	self.introTimer += dt
	if self.introTimer >= 3.0 { // Wait for 3 seconds.
		println("Transitioning to TitleMenuScene...")
		e.SwitchScene("titleMenu")
	}
}

// TextDrawSystem is a DrawSystem for rendering text entities.
type TextDrawSystem struct{}

func (self *TextDrawSystem) Draw(e *katsu2d.Engine, renderer *katsu2d.BatchRenderer) {
	for _, entity := range e.SceneManager().CurrentScene().World.Query(katsu2d.CTText, katsu2d.CTTransform) {
		if tx, ok := e.SceneManager().CurrentScene().World.GetComponent(entity, katsu2d.CTTransform); ok {
			t := tx.(*katsu2d.Transform)
			if txt, ok := e.SceneManager().CurrentScene().World.GetComponent(entity, katsu2d.CTText); ok {
				text := txt.(*katsu2d.Text)
				text.Draw(t.Transform, renderer.GetScreen())
			}
		}
	}
}
