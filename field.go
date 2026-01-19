package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Field [SIZE][SIZE]int

type FieldChange struct {
	changeType string
	from       Position
	to         Position
	value      int
	remove     bool
	merge      bool
}

func (change FieldChange) String() string {
	return fmt.Sprintf("From (%d,%d) to (%d,%d); value: %d; merge: %s; remove: %s", change.from.x, change.from.y, change.to.x, change.to.y, change.value, strconv.FormatBool(change.merge), strconv.FormatBool(change.remove))
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
		newContent := make([]int, SIZE)
		newContentIterator := -1
		localChanges := make([][]Position, SIZE)

		merged := false

		moveCell := func(value int) {
			newContentIterator++

			merged = false
			newContent[newContentIterator] = value

			localChanges[newContentIterator] = append(localChanges[newContentIterator], Position{
				x: iterators["x"],
				y: iterators["y"],
			})
		}

		mergeCell := func(value int) {
			merged = true
			newContent[newContentIterator]++

			localChanges[newContentIterator] = append(localChanges[newContentIterator], Position{x: iterators["x"], y: iterators["y"]})

			// newContentIterator++
		}

		for i := 0; i < SIZE; i++ {
			iterators[internalDirection] = i

			if reverse {
				iterators[internalDirection] = SIZE - i - 1
			}

			value := field[iterators["x"]][iterators["y"]]

			if value == 0 {
				continue
			}

			if newContentIterator == -1 {
				moveCell(value)

				continue
			}

			if newContent[newContentIterator] != value {
				moveCell(value)

				continue
			}

			if merged {
				moveCell(value)

				continue
			}

			mergeCell(value)
		}

		fmt.Println(localChanges)

		for i := 0; i < SIZE; i++ {
			iterators[internalDirection] = i

			if reverse {
				iterators[internalDirection] = SIZE - i - 1
			}

			field[iterators["x"]][iterators["y"]] = newContent[i]

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
	}

	return changes
}

func (field Field) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\n")

	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			fmt.Fprintf(&builder, "%d", field[x][y])
		}
		fmt.Fprintf(&builder, "\n")
	}

	return builder.String()
}

func (field *Field) FromString(str string) {
	str = strings.TrimSpace(str)
	for y, s := range strings.Split(str, "\n") {
		for x, char := range strings.Split(s, " ") {
			value, _ := strconv.Atoi(char)
			field.put(x, y, value)
		}
	}
}
