package main

import (
	"fmt"
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
		}
	}
	return cars
}

type ActiveMessages []ActiveMessage

var (
	instance ActiveMessages
)

func GetMessages() []ActiveMessage {

	if instance == nil {
		instance = make(ActiveMessages, 0)
	}
	return instance
}

func AddMessage(message ActiveMessage) {
	var oldMessages = GetMessages()
	instance = append(oldMessages, message)
}

func updateMessages() {

	for i, f := range GetMessages() {
		if f.currentSize == f.message.warnSize {
			instance = removeMessage(GetMessages(), i)

			fmt.Println("removed message")
			continue
		}
		f.update()
		instance[i] = f
	}
}

func removeMessage(list []ActiveMessage, i int) []ActiveMessage {
	list[i] = list[len(list)-1]
	list[len(list)-1] = ActiveMessage{}
	list = list[:len(list)-1]
	return list
}
