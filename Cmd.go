package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var simulationSeed int64
var simulationWarnSize int
var simulationHops int
var simulationCars int
var simulationObstacles int

func setupFlags() {

	seedPtr := flag.Int64("seed", time.Now().UnixNano(), "int64 simulation seed")
	warnPtr := flag.Int("size", 2, "int warnSize")
	hopPtr := flag.Int("hops", 5, "int hops")
	carsPtr := flag.Int("cars", 30, "int cars")
	obstPtr := flag.Int("obstacles", 1, "int obstacles")

	flag.Parse()

	simulationSeed = *seedPtr
	simulationHops = *hopPtr
	simulationWarnSize = *warnPtr
	simulationCars = *carsPtr
	simulationObstacles = *obstPtr

	rand.Seed(simulationSeed)

	fmt.Println("Seed is: " + strconv.FormatInt(simulationSeed, 10))

}
