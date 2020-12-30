package main

import (
	"github.com/faiface/pixel"
)

type Car struct {
	x, y         float64
	id           string
	sensorActive bool
}

func (car Car) RenderCar(){
	carSprite:= LoadAndSprite("assets/carYellowSmall.png")
	mat := pixel.IM
	mat = pixel.IM.Moved(pixel.V(car.x, car.y))
	carSprite.Draw(mainWindow, mat)
}

func (car *Car) MoveCar(){

}