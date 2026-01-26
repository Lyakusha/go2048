package main

import (
	"time"
)

const SIZE = 4

type Position struct {
	x int
	y int
}

type GameState int

const (
	GameStatePlaying GameState = iota
	GameStateLost
	GameStateWin
)

type Game struct {
	field      Field
	startTime  time.Time
	finishTime time.Time
	score      int
	state      GameState
}

func (game *Game) init() {
	game.field = Field{}
	for range 2 {
		game.field.addRandomTile()
	}

	game.startTime = time.Now()
	game.finishTime = time.Time{}

	game.state = GameStatePlaying
}

func (game *Game) win() {
	game.state = GameStateWin
	game.finishTime = time.Now()
}

func (game *Game) lose() {
	game.state = GameStateLost
	game.finishTime = time.Now()
}

func (game *Game) countScore() int {
	value := 0

	for _, col := range game.field {
		for _, cell := range col {

			if cell != 0 {
				value += 1 << cell
			}
		}
	}

	return value
}

func (game *Game) processMovementAction(action InputAction) (changes []FieldChange) {
	switch action {
	case InputActionMoveUp:
		return game.field.move("x", "y", false)
	case InputActionMoveDown:

		return game.field.move("x", "y", true)
	case InputActionMoveRight:

		return game.field.move("y", "x", true)
	case InputActionMoveLeft:
		return game.field.move("y", "x", false)
	default:
		return nil
	}
}
