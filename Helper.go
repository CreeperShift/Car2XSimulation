package main

import (
	"github.com/faiface/pixel"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"
)

var UP = Move{y: 1}
var DOWN = Move{y: -1}
var LEFT = Move{x: -1}
var RIGHT = Move{x: 1}

var dirALL = []Move{UP, DOWN, LEFT, RIGHT}
var dirUP = []Move{UP, LEFT, RIGHT}
var dirDOWN = []Move{DOWN, LEFT, RIGHT}
var dirLEFT = []Move{UP, LEFT, DOWN}
var dirRIGHT = []Move{UP, DOWN, RIGHT}

var dir = [][]Move{dirUP, dirDOWN, dirLEFT, dirRIGHT}

func compareDir(a, b Move) bool {
	return a.x == b.x && a.y == b.y
}

func IntegerPercentage(x int, f float64) (ret int) {
	return int(math.Round(float64(x) * f))
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func SpriteFromPicture(pic pixel.Picture) *pixel.Sprite {
	return pixel.NewSprite(pic, pic.Bounds())
}

func LoadAndSprite(path string) *pixel.Sprite {
	pic, err := LoadPicture(path)
	if err != nil {
		panic(err)
	}
	return SpriteFromPicture(pic)
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func (m Move) toString() string {

	switch {
	case compareDir(m, UP):
		return "Up"
	case compareDir(m, DOWN):
		return "Down"
	case compareDir(m, LEFT):
		return "Left"
	case compareDir(m, RIGHT):
		return "Right"
	}
	return "Error"
}

func getDistance(x1, y1, x2, y2 float64) float64 {

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

}
