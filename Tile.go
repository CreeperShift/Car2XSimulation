package main

import (
	"github.com/faiface/pixel"
)

type Tile struct {
	x, y         float64
	tileType     int
	rand         int
	obstacle     bool
	obstacleType int
}

var TileSprites = []*pixel.Sprite{LoadAndSprite("assets/TileE.png"), LoadAndSprite("assets/TileStreetPainted.png"), LoadAndSprite("assets/TileStreetPaintedRot.png"), LoadAndSprite("assets/TileStreetN.png")}
var ObstacleSprite = LoadAndSprite("assets/badpix/message-24-warning.png")

func (tile *Tile) setType(i int) {
	tile.tileType = i
}

func (tile Tile) draw() {
	//	var spritesEmpty = []*pixel.Sprite{LoadAndSprite("assets/TileE.png"),LoadAndSprite("assets/TileE2.png"),LoadAndSprite("assets/TileE3.png")}
	mat := pixel.IM
	mat = mat.Moved(pixel.V(tile.x, tile.y))

	/*	if tile.tileType == 0 {
		sprites[0].Draw(mainWindow, mat)
	} else {*/

	TileSprites[tile.tileType].Draw(mainWindow, mat)
	if tile.obstacle {
		ObstacleSprite.Draw(mainWindow, mat)
	}

}
