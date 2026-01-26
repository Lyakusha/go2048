package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const title = "2048"

type Scene struct {
	game         *Game
	tiles        map[Position]*Tile
	leftOffset   int32
	topOffset    int32
	offset       int32
	cellSize     int32
	width        int32
	height       int32
	inputEnabled bool
}

func (scene *Scene) update(dt float32) {
	for _, tile := range scene.tiles {
		go tile.update(dt)
	}
}

func (scene *Scene) init() {
	scene.leftOffset = (scene.width - int32(len(scene.game.field[0]))*scene.cellSize) / 2
	scene.topOffset = (scene.height - int32(len(scene.game.field))*scene.cellSize) / 2
	scene.tiles = scene.buildGrid(scene.game)
}

func (scene *Scene) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	drawFieldBackground(scene.cellSize, scene.offset, scene.leftOffset, scene.topOffset)

	for _, tile := range scene.tiles {
		tile.draw()
	}

	scene.drawTitle()
	scene.drawScore()
	scene.drawTimer()

	rl.EndDrawing()
}

func (scene *Scene) drawTitle() {
	var fontSize int32 = 30
	textSize := rl.MeasureText(title, fontSize)
	rl.DrawText("2048", scene.width/2-textSize/2, 50, fontSize, rl.Black)
}

func (scene *Scene) drawScore() {
	scoreText := strconv.Itoa(scene.game.countScore())
	var fontSize int32 = 20
	textSize := rl.MeasureText(scoreText, fontSize)

	rl.DrawText(scoreText, scene.leftOffset+SIZE*(scene.cellSize+scene.offset)-textSize, 100, fontSize, rl.Black)
}

func (scene *Scene) drawTimer() {
	timer := time.Now().Sub(scene.game.startTime)

	timerText := time.Unix(0, 0).Add(timer).Format("04:05")
	var fontSize int32 = 20

	rl.DrawText(timerText, scene.leftOffset-scene.offset, 100, fontSize, rl.Black)
}

func (scene *Scene) processChanges(changes []FieldChange) {
	waitGroup := sync.WaitGroup{}

	for _, change := range changes {
		if change.from == change.to {
			continue
		}

		tile := scene.tiles[change.from]
		toX := scene.leftOffset + (scene.cellSize+scene.offset)*int32(change.to.x)
		toY := scene.topOffset + (scene.cellSize+scene.offset)*int32(change.to.y)
		duration := 150 * time.Millisecond

		animation := createMoveToAnimation(tile, toX, toY, duration)

		tile.animations[animation] = true

		waitGroup.Go(func() {
			<-animation.stateChannel
			delete(tile.animations, animation)
		})
	}

	waitGroup.Wait()

	fmt.Println("changes processed")

	scene.tiles = scene.buildGrid(scene.game)

	if len(changes) > 0 {
		newTilePosition, newTileValue := scene.game.field.addRandomTile()

		newTile := scene.build(newTilePosition, newTileValue)

		scene.tiles[newTilePosition] = newTile

		animation := createShowTileAnimation(newTile, 150*time.Millisecond)
		newTile.animations[animation] = true

		waitGroup.Go(func() {
			<-animation.stateChannel
			delete(newTile.animations, animation)
		})
	}

	waitGroup.Wait()
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
		x:          x,
		y:          y,
		width:      scene.cellSize,
		height:     scene.cellSize,
		value:      value,
		fontSize:   20,
		animations: make(map[*TileAnimation]bool),
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
