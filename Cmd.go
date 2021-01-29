package main

import (
	"flag"
	"time"
)

var simulationSeed int64
var simulationWarnSize int
var simulationHops int
var simulationCars int
var simulationObstacles int
var simulationTimeCounter int

func setupFlags() {

	seedPtr := flag.Int64("seed", time.Now().UnixNano(), "int64 simulation seed")
	warnPtr := flag.Int("size", 4, "int warnSize")
	hopPtr := flag.Int("hops", 5, "int hops")
	carsPtr := flag.Int("cars", 20, "int cars")
	obstPtr := flag.Int("obstacles", 1, "int obstacles")
	timePtr := flag.Int("time", 1, "int timeCounter")

	flag.Parse()

	simulationSeed = *seedPtr
	simulationHops = *hopPtr
	simulationWarnSize = *warnPtr
	simulationCars = *carsPtr
	simulationObstacles = *obstPtr
	simulationTimeCounter = *timePtr
}
