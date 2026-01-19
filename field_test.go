package main

import (
	"testing"
)

func TestField_FromString(t *testing.T) {
	str := "1 0 0 2\n" +
		"0 0 0 0\n" +
		"0 0 0 0\n" +
		"3 0 0 4\n"

	field := Field{}

	field.FromString(str)

	if field[0][0] != 1 || field[3][0] != 2 || field[0][3] != 3 || field[3][3] != 4 {
		t.Errorf("Field from string fail: \n%s \n %s", str, field)
	}
}

func TestSimpleMoveUp(t *testing.T) {
	field, fieldBeforeMoving := CreateTestField()

	field.move("x", "y", false)

	expectedField := Field{}
	expectedField.FromString("2 2 2 2\n" +
		"0 0 3 2\n" +
		"0 0 0 4\n" +
		"0 0 0 0")

	if !CompareTwoFields(field, expectedField) {
		t.Errorf("Field before moving: \n%s\nField after moving up: %s\nExpected field after moving up: %s", fieldBeforeMoving, field, expectedField)
	}

}

func TestField_MoveDown(t *testing.T) {
	field := Field{}
	fieldBeforeMoving := "2 0 0 0\n" +
		"0 0 0 0\n" +
		"2 0 0 0\n" +
		"4 2 0 0\n"
	field.FromString(fieldBeforeMoving)

	field.move("x", "y", true)

	expectedField := Field{}
	expectedField.FromString("0 0 0 0\n" +
		"0 0 0 0\n" +
		"3 0 0 0\n" +
		"4 2 0 0\n")

	if !CompareTwoFields(field, expectedField) {
		t.Errorf("Field before moving: \n%s\nField after moving down: %s\nExpected field after moving down: %s", fieldBeforeMoving, field, expectedField)
	}
}

func TestField_MoveLeft(t *testing.T) {
	field, fieldBeforeMoving := CreateTestField()

	field.move("y", "x", false)

	expectedField := Field{}
	expectedField.FromString("2 0 0 0\n" +
		"1 3 1 0\n" +
		"3 2 0 0\n" +
		"4 0 0 0")

	if !CompareTwoFields(field, expectedField) {
		t.Errorf("Field before moving: %s\nField after moving left: %s\nExpected field after moving left: %s", fieldBeforeMoving, field, expectedField)
	}

}

func TestField_MoveRight(t *testing.T) {
	field, fieldBeforeMoving := CreateTestField()

	field.move("y", "x", true)

	expectedField := Field{}
	expectedField.FromString("0 0 0 2\n" +
		"0 1 3 1\n" +
		"0 0 3 2\n" +
		"0 0 0 4")

	if !CompareTwoFields(field, expectedField) {
		t.Errorf("Field before moving: %s\nField after moving right: %s\nExpected field after moving right: %s", fieldBeforeMoving, field, expectedField)
	}
}

func CompareTwoFields(test Field, expected Field) bool {
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			if test[x][y] != expected[x][y] {
				return false
			}
		}
	}

	return true
}

func CreateTestField() (Field, string) {
	field := Field{}

	fieldStr := "1 0 0 1\n" +
		"1 2 2 1\n" +
		"0 0 3 2\n" +
		"0 0 0 4"

	field.FromString(fieldStr)

	return field, fieldStr
}
