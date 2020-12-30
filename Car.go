package main

import (
	"github.com/faiface/pixel"
	"math"
	"math/rand"
)

var up = Move{y: 1}
var down = Move{y: -1}
var left = Move{x: -1}
var right = Move{x: 1}

var dirUp = []Move{up, left, right}
var dirDown = []Move{down, left, right}
var dirLeft = []Move{up, left, down}
var dirRight = []Move{up, down, right}

var dir = [][]Move{dirUp, dirDown, dirLeft, dirRight}

type Move struct {
	x, y int
}

type Car struct {
	x, y         int
	id           string
	sensorActive bool
	direction    Move
}

func (car Car) RenderCar() {
	carSprite := LoadAndSprite("assets/carYellowSmall.png")
	mat := pixel.IM
	mat = mat.Moved(pixel.V(streetMap.tiles[car.x][car.y].x, streetMap.tiles[car.x][car.y].y))
	mat = mat.Rotated(pixel.V(streetMap.tiles[car.x][car.y].x, streetMap.tiles[car.x][car.y].y), rotateDirection(car.direction))
	carSprite.Draw(mainWindow, mat)
}

func (car *Car) MoveCar() {

	var movePool []Move

	switch {
	case compareDir(up, car.direction):
		for f := range dir[0] {
			if isInside(*car, f, dir[0]) {
				if streetMap.tiles[car.x+dir[0][f].x][car.y+dir[0][f].y].tileType == 1 {
					movePool = append(movePool, dir[0][f])
				}
			}
		}
	case compareDir(down, car.direction):
		for f := range dir[1] {
			if isInside(*car, f, dir[1]) {
				if streetMap.tiles[car.x+dir[1][f].x][car.y+dir[1][f].y].tileType == 1 {
					movePool = append(movePool, dir[1][f])
				}
			}
		}
	case compareDir(left, car.direction):
		for f := range dir[2] {
			if isInside(*car, f, dir[2]) {
				if streetMap.tiles[car.x+dir[2][f].x][car.y+dir[2][f].y].tileType == 1 {
					movePool = append(movePool, dir[2][f])
				}
			}
		}
	case compareDir(right, car.direction):
		for f := range dir[3] {
			if isInside(*car, f, dir[3]) {
				if streetMap.tiles[car.x+dir[3][f].x][car.y+dir[3][f].y].tileType == 1 {
					movePool = append(movePool, dir[3][f])
				}
			}
		}
	}
	if len(movePool) > 0 {
		i := rand.Intn(len(movePool))
		car.x = car.x + movePool[i].x
		car.y = car.y + movePool[i].y
		car.direction = movePool[i]
	}
}

func compareDir(a, b Move) bool {
	return a.x == b.x && a.y == b.y
}

func isInside(car Car, f int, mov []Move) bool {
	return car.x+mov[f].x <= streetMap.size-1 && car.y+mov[f].y <= streetMap.size-1 && car.x+mov[f].x >= 0 && car.y+mov[f].y >= 0
}

func rotateDirection(m Move) float64 {

	var deg float64

	switch {
	case compareDir(m, up):
		deg = 0
	case compareDir(m, down):
		deg = 180
	case compareDir(m, left):
		deg = 90
	case compareDir(m, right):
		deg = 270
	default:
		deg = 0
	}

	return deg * (math.Pi / 180)
}
