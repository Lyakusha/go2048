package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width int32 = 640
	var height int32 = 480

	rl.InitWindow(width, height, "2048")
	rl.SetTargetFPS(60)

	defer rl.CloseWindow()

	game := Game{}
	game.init()

	inputManager := InputManager{}
	inputManager.init()

	scene := Scene{
		game:     &game,
		offset:   int32(10),
		cellSize: int32(50),
		width:    width,
		height:   height,
	}

	scene.init()

	go func() {
		for {
			action := <-inputManager.outputChannel

			changes := game.processMovementAction(action)

			scene.processChanges(changes)
		}
	}()

	fmt.Println(game.field)

	for !rl.WindowShouldClose() {
		go inputManager.checkInput()

		dt := rl.GetFrameTime()

		scene.update(dt)
		scene.render()
	}
}
