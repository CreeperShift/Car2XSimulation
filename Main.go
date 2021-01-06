package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap

var faster pixel.Rect
var slower pixel.Rect
var speed = 0.5

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

	last := time.Now()
	dt := 0.5
	for !win.Closed() {
		win.Update()
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			buttonPress()
		}
		dt = dt + time.Since(last).Seconds()
		last = time.Now()
		if dt > speed {
			dt = 0.0
			update()
		}
	}
}

func Init() {
	mainWindow.Clear(colornames.Black)
	mainWindow.SetSmooth(true)

	streetMap = NewMap(30, 31)
	streetMap.addStreets()
	streetMap.addCars(10, 30)
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
	var sFaster = LoadAndSprite("assets/bFaster.png")
	var sSlower = LoadAndSprite("assets/bSlower.png")
	mat := pixel.IM
	loc := pixel.V(800, 50)
	mat = mat.Moved(loc)

	sSlower.Draw(mainWindow, mat)
	slower = sSlower.Frame().Moved(loc)
	mat = mat.Moved(pixel.V(45, 0))
	sFaster.Draw(mainWindow, mat)
	faster = sFaster.Frame().Moved(pixel.V(845, 50))

}

func buttonPress() {

	if faster.Contains(mainWindow.MousePosition()) {
		if speed > 0.1 {
			speed -= 0.1
		}
		fmt.Println("press slower")
	}
	if slower.Contains(mainWindow.MousePosition()) {
		speed += 0.1
		fmt.Println("press faster")
	}
}

func main() {
	pixelgl.Run(run)
}
