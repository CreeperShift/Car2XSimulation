package main

import (
	"math/rand"
	"time"
)

type StreetMap struct {
	size     int
	tileSize float64
	tiles    [][]Tile
	cars     []Car
}

func (m *StreetMap) addCars(amount int) {
	rand.Seed(time.Now().UnixNano())
	count := 0

	for x := range m.tiles {
		for y := range m.tiles[x] {
			if m.tiles[x][y].tileType > 0 && count < amount {

				tex := rand.Intn(len(CarSprites))
				car := Car{x: x, y: y, id: "car" + string(rune(count)), sensorActive: false, direction: UP, sprite: CarSprites[tex]}
				m.cars = append(m.cars, car)
				count++
			}
		}
	}

}

func (m *StreetMap) addObstacles(amount, tries int) {
	rand.Seed(time.Now().UnixNano())
	count := 0
	f := 0

	for count < amount {
		count++
		for f < tries {
			f++
			randX := rand.Intn(streetMap.size)
			randY := rand.Intn(streetMap.size)
			if m.tiles[randX][randY].tileType > 0 && !m.tiles[randX][randY].obstacle {
				m.tiles[randX][randY].obstacle = true
				break
			}
		}
	}
}

func (m StreetMap) renderMap() {
	for i := range m.tiles {
		for f := range m.tiles[i] {
			m.tiles[i][f].draw()
		}
	}
}

func NewMap(size int, tilesize float64) *StreetMap {

	s := StreetMap{
		size:     size,
		tileSize: tilesize,
	}
	s.tiles = make([][]Tile, size)
	for i := range s.tiles {
		s.tiles[i] = make([]Tile, size)
	}

	for i := range s.tiles {
		for f := range s.tiles[i] {
			var x, y float64
			x = 17 + float64(i)*s.tileSize
			y = 17 + float64(f)*s.tileSize

			s.tiles[i][f] = Tile{x: x, y: y}
			s.tiles[i][f].rand = rand.Intn(3)
		}
	}
	return &s
}

func (m *StreetMap) addStreets() {

	m.tiles = divideSlice(m.tiles, 2)
	m.tiles = setCorrectStreetTile(m.tiles)
}

func divideSlice(slice [][]Tile, rec int) [][]Tile {
	rand.Seed(time.Now().UnixNano())

	if len(slice) == 0 {
		rec = 0
		return slice
	}

	maxX := rand.Intn(len(slice))
	maxY := rand.Intn(len(slice))

	for x := range slice {
		for y := range slice[x] {

			if x == maxX {
				slice[x][y].setType(1)
			}
			if y == maxY {
				slice[x][y].setType(1)
			}
		}
	}

	if rec > 0 {
		rec--
		rec--
		divideSlice(slice[0:maxX][0:maxY], rec)
	}

	return slice
}

func setCorrectStreetTile(slice [][]Tile) [][]Tile {
	for x := range slice {
		for y := range slice[x] {
			if slice[x][y].tileType > 0 {
				moves := checkDirection(x, y, slice)
				if len(moves) == 1 {
					switch {
					case compareDir(moves[0], RIGHT):
						slice[x][y].tileType = 2
					case compareDir(moves[0], LEFT):
						slice[x][y].tileType = 2
					}
				} else if len(moves) == 2 {
					switch {
					case compareDir(moves[0], UP) && compareDir(moves[1], DOWN):
						slice[x][y].tileType = 1
					case compareDir(moves[0], DOWN) && compareDir(moves[1], UP):
						slice[x][y].tileType = 1
					case compareDir(moves[0], LEFT) && compareDir(moves[1], RIGHT):
						slice[x][y].tileType = 2
					case compareDir(moves[0], RIGHT) && compareDir(moves[1], LEFT):
						slice[x][y].tileType = 2

					}
				} else if len(moves) > 2 {
					slice[x][y].tileType = 3
				}
			}
		}
	}
	return slice
}

func checkDirection(x, y int, slice [][]Tile) (moves []Move) {
	for i := range dirALL {
		if x+dirALL[i].x >= 0 && x+dirALL[i].x < len(slice) && y+dirALL[i].y >= 0 && y+dirALL[i].y < len(slice) {
			//Tile is not outside board
			xNew := x + dirALL[i].x
			yNew := y + dirALL[i].y
			if slice[xNew][yNew].tileType > 0 {
				moves = append(moves, dirALL[i])
			}
		}
	}
	return moves
}

func (m *StreetMap) MoveCars() {
	for f := range m.cars {
		m.cars[f].MoveCar()
	}
}

func (m StreetMap) RenderCars() {
	for i := range m.cars {
		m.cars[i].RenderCar()
	}
}
