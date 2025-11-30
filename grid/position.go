package common

import "fmt"

type Position struct {
	X int
	Y int
}

func (position Position) ToString() string {
	return fmt.Sprintf("%v.%v", position.X, position.Y)
}

func (position Position) WithinBounds(matrix [][]string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return position.X >= 0 && position.X < len(matrix[0]) && position.Y >= 0 && position.Y < len(matrix)
}
