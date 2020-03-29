package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var grassImage *ebiten.Image
var grassTileWidth int
var grassTileHeight int

var tilesInWidth int
var tilesInHeight int

func init() {
	image, _, err := ebitenutil.NewImageFromFile("./assets/grass.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	grassImage = image
	grassTileWidth, grassTileHeight = image.Size()

	tilesInWidth = windowWidth / grassTileWidth
	tilesInHeight = (windowHeight / grassTileHeight) + 1
}

// RenderBackground renders the grass tiles in the background
func RenderBackground(screen *ebiten.Image) {
	for y := 0; y < tilesInHeight; y++ {
		for x := 0; x < tilesInWidth; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*grassTileWidth), float64(y*grassTileHeight))
			_ = screen.DrawImage(grassImage, op)
		}
	}
}
