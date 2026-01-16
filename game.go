package main

const SIZE = 4

type Position struct {
	x int
	y int
}

type Game struct {
	field Field
}

func (game *Game) init() {
	for range 2 {
		game.field.addRandomTile()
	}
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
