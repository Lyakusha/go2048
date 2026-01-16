package main

import (
	"math/rand"
)

type Field [SIZE][SIZE]int

type FieldChange struct {
	from   Position
	to     Position
	value  int
	remove bool
	merge  bool
}

func (field *Field) getEmptyCoordinates() []Position {
	result := make([]Position, 0)

	for x := range field {
		for y := range field[x] {
			if field[x][y] == 0 {
				result = append(result, Position{x: x, y: y})
			}
		}
	}

	return result
}

func (field *Field) put(x int, y int, value int) {
	field[x][y] = value
}

func (field *Field) addRandomTile() (Position, int) {
	emptyCoordinates := field.getEmptyCoordinates()

	randomPosition := emptyCoordinates[rand.Intn(len(emptyCoordinates))]
	newValue := 2

	if rand.Intn(4) == 0 {
		newValue = 3
	}

	field.put(randomPosition.x, randomPosition.y, newValue)

	return randomPosition, newValue
}

func (field *Field) move(externalDirection string, internalDirection string, reverse bool) (changes []FieldChange) {
	iterators := map[string]int{
		"x": 0,
		"y": 0,
	}

	for iterators[externalDirection] = 0; iterators[externalDirection] < len(field); iterators[externalDirection]++ {
		newContent := make([]int, 0)
		localChanges := make([][]Position, 0)

		for i := 0; i < SIZE; i++ {
			iterators[internalDirection] = i

			if reverse {
				iterators[internalDirection] = SIZE - i - 1
			}

			value := field[iterators["x"]][iterators["y"]]

			if value != 0 {
				if len(newContent) > 0 && newContent[len(newContent)-1] == value {
					newContent[len(newContent)-1]++

					if len(localChanges) > 0 {
						localChanges[len(localChanges)-1] = append(localChanges[len(localChanges)-1], Position{x: iterators["x"], y: iterators["y"]})
					} else {
						localChanges = append(localChanges, []Position{
							{x: iterators["x"], y: iterators["y"]},
						})
					}

				} else {
					if i > len(newContent) {
						localChanges = append(localChanges, []Position{
							{x: iterators["x"], y: iterators["y"]},
						})
					}

					newContent = append(newContent, value)
				}
			}
		}

		for i := 0; i < SIZE; i++ {
			iterators[internalDirection] = i

			if reverse {
				iterators[internalDirection] = SIZE - i - 1
			}

			if i < len(newContent) {
				field[iterators["x"]][iterators["y"]] = newContent[i]

				if i < len(localChanges) {
					merge := len(localChanges[i]) > 1
					for changeNumber, from := range localChanges[i] {
						remove := changeNumber > 0

						changes = append(changes, FieldChange{
							from: from,
							to: Position{
								x: iterators["x"],
								y: iterators["y"],
							},
							value:  newContent[i],
							merge:  merge,
							remove: remove,
						})
					}
				}
			} else {
				field[iterators["x"]][iterators["y"]] = 0
			}
		}
	}

	return changes
}
