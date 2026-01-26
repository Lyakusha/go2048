package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var colorsMap = []rl.Color{
	rl.DarkGray,
	rl.LightGray,
	rl.SkyBlue,
	rl.Brown,
	rl.Green,
	rl.DarkGreen,
}

type TileState int

const (
	TileStateIdle TileState = iota
	TileStateMoving
	TileStateMerging
	TileStateAppearing
)

type Tile struct {
	value      int
	x          int32
	y          int32
	width      int32
	height     int32
	animations map[*TileAnimation]bool
	fontSize   int32
}

func (tile *Tile) draw() {
	color := colorsMap[tile.value%len(colorsMap)]

	rl.DrawRectangle(tile.x, tile.y, tile.width, tile.height, color)

	if tile.value != 0 {
		text := fmt.Sprintf("%d", 1<<tile.value)
		textSize := rl.MeasureTextEx(rl.GetFontDefault(), text, float32(tile.fontSize), 0)

		rl.DrawText(text, tile.x+tile.width/2-int32(textSize.X/2), tile.y+tile.height/2-int32(textSize.Y/2), tile.fontSize, rl.Black)
	}
}

func (tile *Tile) update(dt float32) {
	for animation := range tile.animations {
		animation.update(dt)
	}
}
