package main

import (
	"github.com/faiface/pixel"
)

type Car struct {
	x, y         int
	id           string
	sensorActive bool
}

func (car Car) RenderCar() {
	carSprite := LoadAndSprite("assets/carYellowSmall.png")
	mat := pixel.IM
	mat = pixel.IM.Moved(pixel.V(streetMap.tiles[car.x][car.y].x, streetMap.tiles[car.x][car.y].y))
	carSprite.Draw(mainWindow, mat)
}

func (car *Car) MoveCar() {

}
