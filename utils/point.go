package utils

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) ManhattanDistance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func ParsePointFromString(val string) Point {
	// Split the string by the comma
	parts := strings.Split(val, ",")
	if len(parts) != 2 {
		return Point{} // Return zero-value Point if format is invalid
	}

	// Parse X, trimming any accidental spaces
	x, errX := strconv.Atoi(strings.TrimSpace(parts[0]))
	if errX != nil {
		return Point{}
	}

	// Parse Y
	y, errY := strconv.Atoi(strings.TrimSpace(parts[1]))
	if errY != nil {
		return Point{}
	}

	return Point{X: x, Y: y}
}

var (
	Up        = Point{X: 0, Y: -1}
	Down      = Point{X: 0, Y: 1}
	Left      = Point{X: -1, Y: 0}
	Right     = Point{X: 1, Y: 0}
	UpLeft    = Point{-1, -1}
	UpRight   = Point{1, -1}
	DownLeft  = Point{-1, 1}
	DownRight = Point{1, 1}
)

var Directions8 = []Point{
	Up,
	Down,
	Left,
	Right,
	UpLeft,
	UpRight,
	DownLeft,
	DownRight,
}
