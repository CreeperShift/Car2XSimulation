package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Car2Car Simulation",
		Bounds: pixel.R(0, 0, 930, 930),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	mainWindow = win
	Init()
	win.Clear(colornames.Black)

	last := time.Now()
	dt := 0.5
	for !win.Closed() {
		win.Update()
		dt = dt + time.Since(last).Seconds()
		last = time.Now()
		if dt > 0.5 {
			dt = 0.0
			update()
		}
	}
}

func Init() {
	streetMap = NewMap(30, 31)
	streetMap.addStreets()
	streetMap.addCars(10)
	streetMap.addObstacles(3, 15)
}

func update() {
	streetMap.renderMap()
	streetMap.MoveCars()
	streetMap.RenderCars()
	renderButtons()
	//TODO: Check Obstacles
	//TODO: Communication
	//TODO: Warn
	//TODO: Render Comm/Warning
}

func renderButtons() {
	sFaster := LoadAndSprite("assets/bFaster.png")
	sSlower := LoadAndSprite("assets/bSlower.png")
	mat := pixel.IM

	sFaster.Draw(mainWindow, mat)
	sSlower.Draw(mainWindow, mat)

}

func main() {
	pixelgl.Run(run)
}
