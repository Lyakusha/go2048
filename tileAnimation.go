package main

import (
	"time"
)

type TileAnimation struct {
	tile         *Tile
	duration     time.Duration
	rest         time.Duration
	update       func(dt float32)
	stateChannel chan bool
}

func createMoveToAnimation(tile *Tile, toX int32, toY int32, duration time.Duration) *TileAnimation {
	animation := TileAnimation{
		tile:         tile,
		duration:     duration,
		rest:         duration,
		stateChannel: make(chan bool),
	}

	fromX := tile.x
	fromY := tile.y

	animation.update = func(dt float32) {
		animation.rest -= time.Duration(dt * float32(time.Second))

		if animation.rest <= 0 {
			animation.rest = 0
		}

		t := 1 - float32(animation.rest)/float32(animation.duration)

		tile.x = int32(lerp(float32(fromX), float32(toX), t))
		tile.y = int32(lerp(float32(fromY), float32(toY), t))

		if animation.rest == 0 {
			animation.stateChannel <- true
		}
	}

	return &animation
}

func createShowTileAnimation(tile *Tile, duration time.Duration) *TileAnimation {
	animation := TileAnimation{
		tile:         tile,
		duration:     duration,
		rest:         duration,
		stateChannel: make(chan bool),
	}

	fromX := tile.x + tile.width/2
	fromY := tile.y + tile.height/2
	toX := tile.x
	toY := tile.y
	fromWidth := 0
	fromHeight := 0
	toWidth := tile.width
	toHeight := tile.height
	fromFontSize := 1
	toFontSize := tile.fontSize

	animation.update = func(dt float32) {
		animation.rest -= time.Duration(dt * float32(time.Second))

		if animation.rest <= 0 {
			animation.rest = 0
		}

		t := 1 - float32(animation.rest)/float32(animation.duration)

		tile.x = int32(lerp(float32(fromX), float32(toX), t))
		tile.y = int32(lerp(float32(fromY), float32(toY), t))
		tile.width = int32(lerp(float32(fromWidth), float32(toWidth), t))
		tile.height = int32(lerp(float32(fromHeight), float32(toHeight), t))
		tile.fontSize = int32(lerp(float32(fromFontSize), float32(toFontSize), t))

		if animation.rest == 0 {
			animation.stateChannel <- true
		}
	}

	return &animation
}
