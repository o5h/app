package main

import (
	"log"
	"runtime"

	"github.com/o5h/app"
	"github.com/o5h/app/input"
	"github.com/o5h/app/input/mouse"
	"github.com/o5h/glm"
	"github.com/o5h/glx/camera"
	"github.com/o5h/glx/color"
	"github.com/o5h/glx/material"
	"github.com/o5h/glx/mesh/cube"
	"github.com/o5h/glx/renderer"
	_ "github.com/o5h/glx/renderer/gles20"
)

type example struct {
	camera *camera.Camera
	cube   *Object

	targetColor       color.Color
	targetCameraColor color.Color
}

func bindMaterials() {
	renderer.BindMaterial(&material.Default)
}

func (a *example) OnCreate() {
	log.Println("Create")
	a.camera = &camera.Camera{}
	a.camera.SetDefault()

	bindMaterials()
	a.cube = NewObject(cube.New(1.0), material.Default.Instance())
	a.cube.Material.Color = &color.Red
	a.cube.Transform.Scale.Scale(10)

}

func (a *example) OnDestroy() {
	log.Println("Destroy")
}

func (a *example) OnUpdate(delta float64) {
	a.camera.Update()
	a.camera.Color.Lerp(&a.targetCameraColor, float32(delta))
	renderer.SetCamera(a.camera)
	a.cube.Transform.Rotation.X += 1 * float32(delta)
	a.cube.Transform.Rotation.Y += 0.1 * float32(delta)

	a.cube.Material.Color.Lerp(&a.targetColor, float32(delta))
	a.cube.Draw()
}

func (a *example) OnInput(e input.Event) {
	switch i := e.(type) {
	case *mouse.Event:
		if i.Action == mouse.ActionPress && i.Button == mouse.ButtonLeft {
			a.targetColor = color.RandomRGB()
			a.targetCameraColor = color.RandomRGB()
		}
	}

}

func (a *example) OnResize(w, h int) {
	a.camera.SetViewport(glm.RectI{X: 0, Y: 0, W: w, H: h})
}

func init() {
	runtime.LockOSThread()
}

func main() {
	app.Start("Example1", &example{})
}
