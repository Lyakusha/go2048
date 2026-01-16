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

type Tile struct {
	value  int
	x      int32
	y      int32
	width  int32
	height int32
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

func (tile *Tile) moveToAnimation(toX int32, toY int32, duration time.Duration) {
	fromX := tile.x
	fromY := tile.y

	startTime := time.Now()

	ticker := time.NewTicker(100 * time.Millisecond)

	dt := time.Duration(0)

	for dt < duration {
		tickTime := <-ticker.C
		dt = tickTime.Sub(startTime)
		fmt.Println(dt, duration)

		t := float32(dt) / float32(duration)

		tile.x = int32(lerp(float32(fromX), float32(toX), t))
		tile.y = int32(lerp(float32(fromY), float32(toY), t))
	}

	ticker.Stop()
}
