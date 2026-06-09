package engine

type Point struct {
	X, Y int
}

func (p Point) Add(op Point) Point {
	return Point{p.X + op.X, p.Y + op.Y}
}
