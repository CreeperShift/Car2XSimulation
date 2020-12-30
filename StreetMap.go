package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
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

	car := Car{x: 15, y: 15, id: "car1", sensorActive: false, direction: up}
	m.cars = append(m.cars, car)
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

			s.tiles[i][f] = Tile{x, y, 0}
		}
	}
	fmt.Println(s.tiles)
	if test {
		for x := range s.tiles {
			for y := range s.tiles[x] {

				if x == 1 || x == 29 {
					if y >= 1 && y < 30 {
						s.tiles[x][y].setType(1)
					}
				}

				if y == 1 || y == 29 {
					if x >= 1 && x < 30 {
						s.tiles[x][y].setType(1)
					}
				}
			}
		}
	}
	fmt.Println(s.tiles)
	return &s
}

func (tile *Tile) setType(i int) {
	tile.tileType = i
}

func (tile Tile) drawTile() {
	sprites := [2]*pixel.Sprite{LoadAndSprite("assets/TileEmpty.png"), LoadAndSprite("assets/TileStreet.png")}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(tile.x, tile.y))

	sprites[tile.tileType].Draw(mainWindow, mat)
	basicTxt := text.New(pixel.V(tile.x, tile.y), basicAtlas)
	fmt.Fprint(basicTxt, tile.x, " ", tile.y)
	basicTxt.Draw(mainWindow, pixel.IM.Scaled(basicTxt.Orig, 0.9))
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
