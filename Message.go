package main

type Message struct {
	sender      string
	hopCounter  int
	locX, locY  int
	messageCode int
	warnSize    float64
	timeCounter float64
	messageID   uint
}

func NewMessage(car Car) (m *Message) {
	messageIDs++

	m = &Message{}
	m.sender = car.id
	m.locX = car.x
	m.locY = car.y
	m.messageCode = streetMap.tiles[car.x][car.y].obstacleType
	m.hopCounter = simulationHops
	m.warnSize = float64(simulationWarnSize)
	m.messageID = messageIDs

	return m
}

var messageIDs uint = 0

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

func (m Message) isEqual(message Message) bool {
	return m.messageID == message.messageID
}
