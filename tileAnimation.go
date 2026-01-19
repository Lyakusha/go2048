package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileAnimation struct {
	from     rl.Vector2
	to       rl.Vector2
	duration time.Duration
	rest     time.Duration
	update   func(dt float32)
}

func (animation *TileAnimation) Start() {
	animation.rest = animation.duration

	animation.update = func(dt float32) {
	}
}
