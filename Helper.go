package main

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"math"
	"os"
)

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
