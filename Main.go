package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap

var last time.Time
var dt float64

func main() {
	pixelgl.Run(run)
}

func run() {
	setupWindow()
	Init()
	for !mainWindow.Closed() {
		handleButtons()
		simulate()
		mainWindow.Update()
	}
}

func setupWindow() {
	cfg := pixelgl.WindowConfig{
		Title:  "Car2Car Hinderniswarnung Simulation",
		Bounds: pixel.R(0, 0, sizeX, sizeY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	mainWindow = win
}

func simulate() {

	dt = dt + time.Since(last).Seconds()
	last = time.Now()
	if dt > speed {
		dt = 0.0
		if len(GetMessages()) == 0 {
			update()
		} else {
			updateWifi()
		}
	}
}

func updateWifi() {
	mainWindow.Clear(colornames.Lightgray)
	streetMap.renderMap()
	streetMap.RenderCars()
	updateMessages()
	renderButtons()
	renderGUI()
}

func Init() {
	setupFlags()
	SetupFonts()
	guiBase = pixel.V(930, 0)
	mainWindow.Clear(colornames.Lightgray)
	mainWindow.SetSmooth(true)
	streetMap = NewMap(30, 30.5)
	streetMap.init()
	last = time.Now()
	dt = 0.5
}

func update() {
	mainWindow.Clear(colornames.Lightgray)
	streetMap.renderMap()
	streetMap.UpdateCars()
	streetMap.RenderCars()
	renderButtons()
	renderGUI()
}
