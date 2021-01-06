package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap
var basicAtlas *text.Atlas

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Car2Car Simulation",
		Bounds: pixel.R(0, 0, 930, 930),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	mainWindow = win
	basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	Init()
	win.Clear(colornames.Black)

	for !win.Closed() {
		win.Update()
		update()
	}
}

func Init() {
	streetMap = NewMap(30, 31)
	streetMap.addStreets()
	streetMap.addCar()
	//TODO: spawn obstacle (1 or more)
}

func update() {
	streetMap.renderMap()
	streetMap.MoveCars()
	streetMap.RenderCars()
	//TODO: Check Obstacles
	//TODO: Communication
	//TODO: Warn
	//TODO: Render Comm/Warning
}

func main() {
	pixelgl.Run(run)
}
