package app

import (
	"time"

	"github.com/o5h/app/input"
	"github.com/o5h/egl"
)

type context struct {
	Done           bool
	NativeWindow   egl.NativeWindow
	NativeDisplay  egl.NativeDisplay
	Context        egl.Context
	Surface        egl.Surface
	Display        egl.Display
	App            App
	LastUpdateTime time.Time
}

func (ctx *context) onCreate() {
	ctx.App.OnCreate()
	ctx.LastUpdateTime = time.Now()
}

func (ctx *context) onClose() {
	ctx.Done = true
}

func (ctx *context) onDestroy() {
	ctx.App.OnDestroy()
}

func (ctx *context) onResize(w, h int) {
	ctx.App.OnResize(w, h)
}

func (ctx *context) onUpdate(now float64) {
	ctx.App.OnUpdate(now)
	egl.SwapBuffers(ctx.Display, ctx.Surface)
}

func (ctx *context) onInput(e input.Event) {
	ctx.App.OnInput(e)
}
