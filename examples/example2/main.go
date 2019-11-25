package main

import (
	"log"
	"os"
	"runtime"

	"github.com/o5h/app"
	"github.com/o5h/app/input"
	"github.com/o5h/app/input/mouse"
	"github.com/o5h/gles20"
	"github.com/o5h/openal/al"
	"github.com/o5h/openal/alc"
	"github.com/o5h/wav"
)

type example struct {
	device  alc.Device
	sources []uint32
	buffers []uint32
}

func (app *example) OnCreate() {
	log.Println("Create")

	file, err := os.Open("Alesis-Sanctuary-QCard-Tines-Aahs-C4.wav")
	checkErr(err)
	defer file.Close()
	w := &wav.Wav{}
	w.ReadFrom(file)

	//initialize audio
	app.device = alc.OpenDevice("")
	context := alc.CreateContext(app.device, nil)
	alc.MakeContextCurrent(context)

	app.sources = []uint32{0}
	al.GenSources(app.sources)
	al.Sourcef(app.sources[0], al.AL_PITCH, 1)
	al.Sourcef(app.sources[0], al.AL_GAIN, 1)
	al.Source3f(app.sources[0], al.AL_POSITION, 0, 0, 0)
	al.Source3f(app.sources[0], al.AL_VELOCITY, 0, 0, 0)
	al.Sourcei(app.sources[0], al.AL_LOOPING, 0)

	app.buffers = []uint32{0}
	al.GenBuffers(app.buffers)

	al.BufferData(app.buffers[0], al.AL_FORMAT_STEREO16, w.Data, w.SampleRate)
	al.Sourcei(app.sources[0], al.AL_BUFFER, int32(app.buffers[0]))
}

func (app *example) OnDestroy() {
	log.Println("Destroy")
	al.DeleteBuffers(app.buffers)
	al.DeleteSources(app.sources)
	alc.CloseDevice(app.device)
}

func (a *example) OnUpdate(delta float64) {
	gles20.Clear(gles20.COLOR_BUFFER_BIT | gles20.DEPTH_BUFFER_BIT | gles20.STENCIL_BUFFER_BIT)
	gles20.ClearColor(0.5, 0.5, 1.0, 0)
}

func (app *example) OnInput(e input.Event) {
	switch i := e.(type) {
	case *mouse.Event:
		if i.Action == mouse.ActionPress && i.Button == mouse.ButtonLeft {
			al.SourcePlay(app.sources[0])
		}
	}
}

func (a *example) OnResize(w, h int) {
	gles20.Viewport(0, 0, w, h)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	app.Start("Example2", &example{})
}
