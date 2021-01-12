package main

import (
	"fmt"
	"github.com/faiface/pixel"
)

type ActiveMessage struct {
	currentSize float64
	locX, locY  float64
	message     Message
}

func NewActiveMessage(car Car, message Message) *ActiveMessage {
	return &ActiveMessage{locX: float64(car.x), locY: float64(car.y), message: message, currentSize: 1}
}

func (m *ActiveMessage) update() {

	cars := getCarsInArea(m.locX, m.locY, m.currentSize)

	for _, c := range cars {
		c.receiveMessage(m.message)
	}
	m.currentSize++
}

func getCarsInArea(locX, locY float64, size float64) (cars []*Car) {

	var ico = LoadAndSprite("assets/badpix/message-16-info.png")

	lowerBoundX := locX - size
	lowerBoundX = ClampValue(lowerBoundX, 0)
	lowerBoundY := locY - size
	lowerBoundY = ClampValue(lowerBoundY, 0)

	upperBoundX := locX + size
	upperBoundX = ClampValue(upperBoundX, float64(len(streetMap.tiles)-1))

	upperBoundY := locY + size
	upperBoundY = ClampValue(upperBoundY, float64(len(streetMap.tiles)-1))

	for i := lowerBoundX; i <= upperBoundX; i++ {
		for f := lowerBoundY; f <= upperBoundY; f++ {
			hasCar, car := streetMap.getCarByLocation(int(i), int(f))
			if hasCar {
				cars = append(cars, car)
			}
			mat := pixel.IM

			ii := int(i)
			ff := int(f)

			v := pixel.V(streetMap.tiles[ii][ff].x, streetMap.tiles[ii][ff].y)
			mat = mat.Moved(v)
			ico.Draw(mainWindow, mat)
		}
	}
	return cars
}

type ActiveMessages []*ActiveMessage

var (
	instance       ActiveMessages
	futureInstance ActiveMessages
)

func GetMessages() []*ActiveMessage {

	if instance == nil {
		instance = make(ActiveMessages, 0)
	}
	return instance
}

func AddMessage(message ActiveMessage) {
	var oldMessages = GetMessages()
	instance = append(oldMessages, &message)
}

func updateMessages() {

	instance = cleanInstance()

	for _, f := range instance {
		f.update()
	}

	fmt.Println(futureInstance)

	instance = append(instance, futureInstance...)
	futureInstance = make(ActiveMessages, 0)
}

func cleanInstance() ActiveMessages {
	cleanedInstance := make(ActiveMessages, 0)

	for _, f := range instance {
		if f.currentSize < f.message.warnSize {
			cleanedInstance = append(cleanedInstance, f)
		}
	}
	return cleanedInstance
}
