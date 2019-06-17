package app

import (
	"github.com/o5h/app/input"
)

type App interface {
	OnCreate()
	OnDestroy()
	OnUpdate(now float64)
	OnResize(w, h int)
	OnInput(input.Event)
}

func Start(title string, app App) {
	context := create(title, app)
	defer context.onDestroy()
	context.mainLoop()
}
