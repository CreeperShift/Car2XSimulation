package main

type Message struct {
	sender           string
	hopCounter       int
	locX, locY       float64
	senderX, senderY float64
	messageCode      int
	warnSize         float64
	timeCounter      float64
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
