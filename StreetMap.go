package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"math/rand"
	"time"
)

type (
	Tile struct {
		x, y     float64
		tileType int
	}
	Obstacle struct {
		x, y int
	}
	StreetMap struct {
		size      int
		tiles     [][]Tile
		cars      []Car
		obstacles []Obstacle
	}
)

func (m *StreetMap) addCar() {

	//TODO: proper car spawning

	sprites := []*pixel.Sprite{LoadAndSprite("assets/car1.png"), LoadAndSprite("assets/car2.png"), LoadAndSprite("assets/car3.png"), LoadAndSprite("assets/car4.png")}

	rand.Seed(time.Now().UnixNano())
	count := 0

	for x := range m.tiles {
		for y := range m.tiles[x] {
			if m.tiles[x][y].tileType == 1 && count < 10 {
				tex := rand.Intn(len(sprites))
				car := Car{x: x, y: y, id: "car1", sensorActive: false, direction: up, sprite: sprites[tex]}
				m.cars = append(m.cars, car)

				count++
			}
		}
	}

}

func (m StreetMap) renderMap() {
	for i := range m.tiles {
		for f := range m.tiles[i] {
			m.tiles[i][f].drawTile()
		}
	}
}

func NewMap(size int, test bool) *StreetMap {

	s := StreetMap{
		size: size,
	}
	s.tiles = make([][]Tile, size)
	for i := range s.tiles {
		s.tiles[i] = make([]Tile, size)
	}

	fmt.Println(s.tiles)

	for i := range s.tiles {
		for f := range s.tiles[i] {
			var x, y float64
			x = 15 + float64(i)*30
			y = 15 + float64(f)*30

			s.tiles[i][f] = Tile{x: x, y: y}
		}
	}
	fmt.Println(s.tiles)
	return &s
}

func (m *StreetMap) addStreets() {

	m.tiles = divideSlice(m.tiles, 2)
}

func divideSlice(slice [][]Tile, rec int) [][]Tile {
	rand.Seed(time.Now().UnixNano())
	maxX := rand.Intn(len(slice) - 1)
	maxY := rand.Intn(len(slice) - 1)

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
		divideSlice(slice[0:maxX][0:maxY], rec)
	}

	return slice
}

func (tile *Tile) setType(i int) {
	tile.tileType = i
}

func (tile Tile) drawTile() {
	sprites := [2]*pixel.Sprite{LoadAndSprite("assets/TileEmpty.png"), LoadAndSprite("assets/TileStreet.png")}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(tile.x, tile.y))

	sprites[tile.tileType].Draw(mainWindow, mat)
	/*		basicTxt := text.New(pixel.V(tile.x, tile.y), basicAtlas)
			fmt.Fprintln(basicTxt, tile.x, " ", tile.y)
			basicTxt.Draw(mainWindow, pixel.IM.Scaled(basicTxt.Orig, 0.5))*/
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
