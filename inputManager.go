package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputAction int

const (
	InputActionMoveUp InputAction = iota
	InputActionMoveRight
	InputActionMoveDown
	InputActionMoveLeft
)

type InputManager struct {
	keyMapping    map[int32]InputAction
	outputChannel chan InputAction
}

func (inputManager *InputManager) init() {
	inputManager.outputChannel = make(chan InputAction, 1)

	inputManager.keyMapping = map[int32]InputAction{
		rl.KeyW: InputActionMoveUp,
		rl.KeyD: InputActionMoveRight,
		rl.KeyS: InputActionMoveDown,
		rl.KeyA: InputActionMoveLeft,
	}
	inputManager.keyMapping[rl.KeyW] = InputActionMoveUp
	inputManager.keyMapping[rl.KeyD] = InputActionMoveRight
	inputManager.keyMapping[rl.KeyS] = InputActionMoveDown
	inputManager.keyMapping[rl.KeyA] = InputActionMoveLeft
}

func (inputManager *InputManager) checkInput() {
	for key, value := range inputManager.keyMapping {
		if rl.IsKeyReleased(key) {
			inputManager.outputChannel <- value
		}
	}
}
