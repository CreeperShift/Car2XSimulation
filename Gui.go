package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"strconv"
)

var sizeX float64 = 1330
var sizeY float64 = 930

var faster pixel.Rect
var slower pixel.Rect
var speed = 0.5
var guiBase pixel.Vec

var FontHeader *text.Atlas
var FontText *text.Atlas

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
