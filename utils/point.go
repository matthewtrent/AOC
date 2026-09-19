package utils

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

var (
	Up    = Point{X: 0, Y: -1}
	Down  = Point{X: 0, Y: 1}
	Left  = Point{X: -1, Y: 0}
	Right = Point{X: 1, Y: 0}
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
