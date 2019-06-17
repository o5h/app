package main

import (
	"log"
	"runtime"

	"github.com/o5h/app"
	"github.com/o5h/app/input"
	"github.com/o5h/gles20"
)

type example struct {
}

func (a *example) OnCreate() {
	log.Println("Create")
}

func (a *example) OnDestroy() {
	log.Println("Destroy")
}

func (a *example) OnUpdate(delta float64) {
	gles20.Clear(gles20.COLOR_BUFFER_BIT | gles20.DEPTH_BUFFER_BIT | gles20.STENCIL_BUFFER_BIT)
	gles20.ClearColor(0.0, 1.0, 1.0, 0)

}

func (a *example) OnInput(e input.Event) {
	log.Println(e)
}

func (a *example) OnResize(w, h int) {
	gles20.Viewport(0, 0, w, h)
}

func init() {
	runtime.LockOSThread()
}

func main() {
	app.Start("Example1", &example{})
}
