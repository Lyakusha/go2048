package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var colorsMap = []rl.Color{
	rl.Gray,
	rl.DarkGray,
	rl.LightGray,
	rl.SkyBlue,
	rl.Brown,
	rl.Green,
	rl.DarkGreen,
}

func noUpdate(dt float32) {}

type Tile struct {
	value  int
	x      int32
	y      int32
	width  int32
	height int32
	update func(dt float32)
}

func (tile *Tile) draw() {
	color := colorsMap[tile.value%len(colorsMap)]

	rl.DrawRectangle(tile.x, tile.y, tile.width, tile.height, color)

	if tile.value != 0 {
		text := fmt.Sprintf("%d", int(math.Pow(2, float64(tile.value-1))))
		textSize := rl.MeasureTextEx(rl.GetFontDefault(), text, 20, 0)

		rl.DrawText(text, tile.x+tile.width/2-int32(textSize.X/2), tile.y+tile.height/2-int32(textSize.Y/2), 20, rl.Black)
	}
}

func (tile *Tile) moveToAnimation(toX int32, toY int32, duration time.Duration) func(dt float32) {
	rest := duration

	fromX := tile.x
	fromY := tile.y

	update := func(dt float32) {
		if rest < 0 {
			return
		}

		t := dt / float32(duration)

		tile.x = int32(lerp(float32(fromX), float32(toX), t))
		tile.y = int32(lerp(float32(fromY), float32(toY), t))

		rest -= time.Duration(dt * float32(time.Second))

		if rest < 0 {
			rest = 0
		}
	}

	return update
}
