package main

type Point struct {
	x, y int
}

func (p Point) Add(op Point) Point {
	return Point{p.x + op.x, p.y + op.y}
}
