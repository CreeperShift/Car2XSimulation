package main

import (
	"github.com/faiface/pixel"
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
	car := Car{x: 15, y: 15, id: "car1", sensorActive: false}
	m.cars = append(m.cars, car)
}

func (m StreetMap) renderMap() {
	for i := range m.tiles {
		for f := range m.tiles[i] {
			m.tiles[i][f].drawTile()
		}
	}
}

func NewMap(size int) *StreetMap {

	s := StreetMap{
		size: size,
	}
	s.tiles = make([][]Tile, size)
	for i := range s.tiles {
		s.tiles[i] = make([]Tile, size)
	}
	for i := range s.tiles {
		for f := range s.tiles[i] {
			var x, y float64
			x = 15 + float64(i)*30
			y = 15 + float64(f)*30

			s.tiles[i][f] = Tile{x, y, 0}
		}
	}

	return &s
}

func (tile Tile) drawTile() {
	sprites := [2]*pixel.Sprite{LoadAndSprite("assets/TileEmpty.png"), LoadAndSprite("assets/TileStreet.png")}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(tile.x, tile.y))

	sprites[tile.tileType].Draw(mainWindow, mat)
}

func (m StreetMap) RenderCars() {
	for i := range m.cars {
		m.cars[i].RenderCar()
	}
}
