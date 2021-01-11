package main

import (
	"github.com/faiface/pixel"
)

type ActiveMessage struct {
	currentSize float64
	locX, locY  float64
	message     Message
	messageID   uint
}

var activeMessageIDs uint = 0

func NewActiveMessage(car Car, message Message) *ActiveMessage {
	activeMessageIDs++
	return &ActiveMessage{locX: float64(car.x), locY: float64(car.y), message: message, currentSize: 1, messageID: activeMessageIDs}
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
	//	fmt.Println("Loc was: ", locX, ", ", locY, " clamped to: ", lowerBoundX, " , ", upperBoundX, " - ", lowerBoundY, ", ", upperBoundX)

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

type ActiveMessages []ActiveMessage

var (
	instance ActiveMessages
	queue    ActiveMessages
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

	var uInstance = make(ActiveMessages, 0)

	for _, f := range GetMessages() {
		if f.currentSize < f.message.warnSize {
			uInstance = append(uInstance, f)
		}
	}
	instance = uInstance
	queue = uInstance
	if len(queue) > 0 {
		updateMessageRec()
	}
}

func (m *ActiveMessages) getActiveMessageFromID(id uint) int {
	for i, e := range *m {
		if e.messageID == id {
			return i
		}
	}
	return 0
}

func updateMessageRec() {
	m := queue[0]
	m.update()
	i := instance.getActiveMessageFromID(m.messageID)
	instance[i].currentSize = m.currentSize
	queue = queue[1:]
	if len(queue) > 0 {
		updateMessageRec()
	}

	//TODO: Fix Followup shit

}
