package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"strconv"
	"time"
)

var mainWindow *pixelgl.Window
var streetMap *StreetMap

var sizeX float64 = 1330
var sizeY float64 = 930

var faster pixel.Rect
var slower pixel.Rect
var speed = 0.5
var guiBase pixel.Vec

//Fonts
var FontHeader *text.Atlas
var FontText *text.Atlas

var last time.Time
var dt float64

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Car2Car Simulation",
		Bounds: pixel.R(0, 0, sizeX, sizeY),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	mainWindow = win
	Init()

	last = time.Now()
	dt = 0.5
	for !win.Closed() {

		handleButtons()
		handleSimulation()

		//TODO: LERP CARS
		win.Update()
	}
}

func handleSimulation() {

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
	updateMessages()
}

func handleButtons() {
	if mainWindow.JustPressed(pixelgl.MouseButtonLeft) {
		buttonPress()
	}
}

func SetupFonts() {
	header, err := loadTTF("assets/font/Poppins-Regular.ttf", 22)
	if err != nil {
		panic(err)
	}
	textF, err := loadTTF("assets/font/Roboto-Regular.ttf", 14)
	if err != nil {
		panic(err)
	}
	FontHeader = text.NewAtlas(header, text.ASCII)
	FontText = text.NewAtlas(textF, text.ASCII)
}

func Init() {
	SetupFonts()
	guiBase = pixel.V(930, 0)
	mainWindow.Clear(colornames.Lightgray)
	mainWindow.SetSmooth(true)
	streetMap = NewMap(30, 30.5)
	streetMap.addStreets()
	streetMap.addObstacles(3, 50)
	streetMap.addCars(20, 300)
}

func update() {
	mainWindow.Clear(colornames.Lightgray)
	streetMap.renderMap()
	streetMap.UpdateCars()
	streetMap.RenderCars()
	renderButtons()
	renderGUI()
	//TODO: Check Obstacles
	//TODO: Communication
	//TODO: Warn
	//TODO: Render Comm/Warning
}

func renderGUI() {
	//var spacing float64 = 20
	var center = guiBase.Add(pixel.V(160, 900))
	basicTxt := text.New(center, FontHeader)
	basicTxt.Color = colornames.Black
	fmt.Fprintln(basicTxt, "Cars")
	basicTxt.Draw(mainWindow, pixel.IM)

	carTxt := text.New(center.Add(pixel.V(-150, -30)), FontText)
	carTxt.LineHeight = FontText.LineHeight() * 1.5
	carTxt.Color = colornames.Black

	var carIDs = []string{""}
	for _, line := range streetMap.cars {

		txt := line.id + ": [X:" + strconv.FormatInt(int64(line.x), 10) + "|Y:" + strconv.FormatInt(int64(line.y), 10) + "] "
		txt = txt + "Direction: " + line.direction.toString()

		carIDs = append(carIDs, txt)
	}

	for _, line := range carIDs {
		fmt.Fprintln(carTxt, line)
	}

	carTxt.Draw(mainWindow, pixel.IM)

}

func renderButtons() {
	var sFaster = LoadAndSprite("assets/bFaster.png")
	var sSlower = LoadAndSprite("assets/bSlower.png")
	mat := pixel.IM
	loc := pixel.V(1000, 50)
	mat = mat.Moved(loc)

	sSlower.Draw(mainWindow, mat)
	slower = sSlower.Frame().Moved(loc)
	mat = mat.Moved(pixel.V(45, 0))
	sFaster.Draw(mainWindow, mat)
	faster = sFaster.Frame().Moved(pixel.V(1045, 50))

	basicTxt := text.New(loc.Add(pixel.V(70, -7)), FontText)
	basicTxt.Color = colornames.Black
	fmt.Fprintln(basicTxt, "Delay: "+strconv.FormatFloat(speed, 'f', 1, 32)+"s.")
	basicTxt.Draw(mainWindow, pixel.IM)

}

func buttonPress() {

	if faster.Contains(mainWindow.MousePosition()) {
		if speed > 0.1 {
			speed -= 0.1
		}
		fmt.Println("Faster, delay: " + strconv.FormatFloat(speed, 'f', 1, 32) + "s.")
	}
	if slower.Contains(mainWindow.MousePosition()) {
		speed += 0.1
		fmt.Println("Slower, delay: " + strconv.FormatFloat(speed, 'f', 1, 32) + "s.")
	}
}

func main() {
	pixelgl.Run(run)
}
