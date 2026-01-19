package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	game       *Game
	actions    map[*Tile]func(dt float32)
	tiles      map[Position]*Tile
	leftOffset int32
	topOffset  int32
	offset     int32
	cellSize   int32
	width      int32
	height     int32
}

func (scene *Scene) update(dt float32) {
	for _, tile := range scene.tiles {
		if tile.update != nil {
			tile.update(dt)
		}
	}

	for _, action := range scene.actions {
		action(dt)
	}
}

func (scene *Scene) init() {
	scene.leftOffset = (scene.width - int32(len(scene.game.field[0]))*scene.cellSize) / 2
	scene.topOffset = (scene.height - int32(len(scene.game.field))*scene.cellSize) / 2
	scene.tiles = scene.buildGrid(scene.game)
	scene.actions = make(map[*Tile]func(dt float32))
}

func (scene *Scene) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	drawFieldBackground(scene.cellSize, scene.offset, scene.leftOffset, scene.topOffset)

	for _, tile := range scene.tiles {
		tile.draw()
	}

	rl.EndDrawing()
}

func (scene *Scene) processChanges(changes []FieldChange) {
	for _, change := range changes {
		fmt.Println(change)

		tile := scene.tiles[change.from]

		fromX := tile.x
		fromY := tile.y
		toX := scene.leftOffset + (scene.cellSize+scene.offset)*int32(change.to.x)
		toY := scene.topOffset + (scene.cellSize+scene.offset)*int32(change.to.y)
		duration := 250 * time.Millisecond
		rest := duration

		scene.actions[tile] = func(dt float32) {
			rest -= time.Duration(dt * float32(time.Second))

			t := 1 - float32(rest)/float32(duration)

			if t > 1 {
				t = 1
				delete(scene.actions, tile)

				delete(scene.tiles, change.from)

				if !change.remove {
					scene.tiles[change.to] = tile
				}

				if change.merge {
					tile.value = change.value
				}

				fmt.Println(scene.tiles)
			}

			tile.x = int32((1-t)*float32(fromX) + t*float32(toX))
			tile.y = int32((1-t)*float32(fromY) + t*float32(toY))

			fmt.Println(t)
		}

		//waitGroup.Go(func() {
		//		action := tile.moveToAnimation(scene.leftOffset+(scene.cellSize+scene.offset)*int32(change.to.x), scene.topOffset+(scene.cellSize+scene.offset)*int32(change.to.y), 2000*time.Millisecond)
		//			scene.actions = append(scene.actions, action)
		//
		//			if change.merge {
		//				tile.value = change.value
		//			}
		//		})
	}

	//scene.tiles = scene.buildGrid(scene.game)

	if len(changes) > 0 {
		newTilePosition, newTileValue := scene.game.field.addRandomTile()

		scene.tiles[newTilePosition] = scene.build(newTilePosition, newTileValue)
	}
}

func drawFieldBackground(cellSize int32, offset int32, leftOffset int32, topOffset int32) {
	rl.DrawRectangle(leftOffset-offset, topOffset-offset, (cellSize+offset)*SIZE+offset, (cellSize+offset)*SIZE+offset, rl.Gray)
}

func (scene *Scene) getCoordinatesForPositionInGrid(x int, y int) (int32, int32) {
	return scene.leftOffset + (scene.cellSize+scene.offset)*int32(x),
		scene.topOffset + (scene.cellSize+scene.offset)*int32(y)
}

func (scene *Scene) build(pos Position, value int) *Tile {
	x, y := scene.getCoordinatesForPositionInGrid(pos.x, pos.y)

	return &Tile{
		x:      x,
		y:      y,
		width:  scene.cellSize,
		height: scene.cellSize,
		value:  value,
	}
}

func (scene *Scene) buildGrid(game *Game) map[Position]*Tile {
	grid := map[Position]*Tile{}

	for x, column := range game.field {
		for y := range column {
			value := game.field[x][y]

			if value > 0 {
				pos := Position{x, y}
				grid[pos] = scene.build(pos, value)
			}
		}
	}

	return grid
}
