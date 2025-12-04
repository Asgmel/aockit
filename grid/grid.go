package grid

import "fmt"

type Position struct {
	X int
	Y int
}

func (position Position) ToString() string {
	return fmt.Sprintf("%v.%v", position.X, position.Y)
}

func WithinBounds[T any](position Position, matrix [][]T) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return position.X >= 0 && position.X < len(matrix[0]) && position.Y >= 0 && position.Y < len(matrix)
}

func GetNeighbouringPositions[T any](position Position, matrix [][]T) []Position {
	neighbours := []Position{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			currentPos := Position{X: position.X + x, Y: position.Y + y}
			if WithinBounds(currentPos, matrix) {
				neighbours = append(neighbours, currentPos)
			}
		}
	}
	return neighbours
}
