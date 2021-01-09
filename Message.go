package main

type Message struct {
	sender      string
	hopCounter  int
	locX, locY  int
	messageCode int
	warnSize    float64
	timeCounter float64
}

func (message *Message) HopCounter() int {
	return message.hopCounter
}

func (message *Message) SetHopCounter(hopCounter int) {
	message.hopCounter = hopCounter
}

var messageCodes = []int{
	2, 3, 12, 91, 92, 97,
}

var messageCodeMapped = map[int]string{
	2:  "accident",
	3:  "roadwork",
	12: "humanPresenceOnTheRoad",
	91: "vehicleBreakdown",
	92: "postCrash",
	97: "collisionRisk",
}

func (message *Message) setSender(s string) {
	message.sender = s
}

func (message *Message) setHops(i int) {
	message.hopCounter = i
}
func (message *Message) setXY(x, y int) {
	message.locX = x
	message.locY = y
}
func (message *Message) setCode(i int) {
	message.messageCode = i
}
func (message *Message) setWarnSize(i float64) {
	message.warnSize = i
}
func (message *Message) setTimeCounter(i float64) {
	message.timeCounter = i
}

func sendMessage(car *Car, message Message) {
	for _, c := range streetMap.cars {

		distanceWifi := getDistance(float64(car.x), float64(car.y), float64(c.x), float64(c.y))
		if distanceWifi < WifiDistance {

			distance := getDistance(float64(c.x), float64(c.y), float64(message.locX), float64(message.locY))
			if distance < message.warnSize {
				streetMap.sendMessageToCar(c.id, message)
			}
		}

	}
}

func createMessage(car Car) Message {
	m := Message{}
	m.sender = car.id
	m.locX = car.x
	m.locY = car.y
	m.messageCode = streetMap.tiles[car.x][car.y].obstacleType
	m.hopCounter = 3
	m.timeCounter = 5
	m.warnSize = 15
	return m
}
