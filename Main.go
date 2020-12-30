package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Car2Car Simulation",
		Bounds: pixel.R(0, 0, 900, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	mainWindow = win

	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		win.Update()
		update()
	}
}


func init(){
	streetMap = NewMap(30)
	streetMap.addCar()
	//TODO: Spawn cars
}

func update() {
	streetMap.renderMap()
	//TODO: Move cars
	streetMap.RenderCars()
	//TODO: Check Obstacles
	//TODO: Communication
	//TODO: Warn
	//TODO: Render Comm/Warning
}


func main() {
	pixelgl.Run(run)
}
