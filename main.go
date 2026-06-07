package main

import (
	"game-of-life/graphics"
	"image/color"
	"runtime"
)

const width = 20
const height = 20

type Point struct {
	x, y int
}

func (p Point) Add(op Point) Point {
	return Point{p.x + op.x, p.y + op.y}
}

type Game struct {
	state     [][]bool
	nextState [][]bool
}

func (g *Game) At(p Point) bool {
	return g.state[p.y][p.x]
}

func (g *Game) SetNext(p Point, alive bool) {
	g.nextState[p.y][p.x] = alive
}

func initGame(w, h int) *Game {
	state := make([][]bool, h)
	nextState := make([][]bool, h)
	for i := range state {
		state[i] = make([]bool, w)
		nextState[i] = make([]bool, w)
	}
	return &Game{state, nextState}
}

func (g *Game) calculateNextState() {
	for y := 0; y < len(g.state); y++ {
		for x := 0; x < len(g.state[0]); x++ {
			g.calculateNextCell(Point{x, y})
		}
	}
}

func (g *Game) calculateNextCell(p Point) {
	dirs := []Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	nCount := 0
	for _, dir := range dirs {
		n := p.Add(dir)
		yCheck := n.y >= 0 && n.y < len(g.state)
		xCheck := n.x >= 0 && n.x < len(g.state[0])
		if yCheck && xCheck && g.At(n) {
			nCount++
		}
	}
	if g.At(p) {
		if nCount < 2 {
			g.SetNext(p, false)
		} else if nCount > 3 {
			g.SetNext(p, false)
		} else {
			g.SetNext(p, true)
		}
	} else {
		g.SetNext(p, nCount == 3)
	}
}

func (g *Game) clearNextState() {
	for y := range g.nextState {
		for x := range g.nextState[y] {
			g.nextState[y][x] = false
		}
	}
}

func (g *Game) Swap() {
	g.state, g.nextState = g.nextState, g.state
	g.clearNextState()
}

func (g *Game) SeedGlider(origin Point) {
	// Glider shape relative to origin:
	// . X .
	// . . X
	// X X X
	offsets := []Point{
		{1, 0},
		{2, 1},
		{0, 2}, {1, 2}, {2, 2},
	}
	for _, offset := range offsets {
		p := origin.Add(offset)
		if p.y < len(g.state) && p.x < len(g.state[0]) {
			g.state[p.y][p.x] = true
		}
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	game := initGame(width, height)

	game.SeedGlider(Point{2, 2})

	cellSize := 20

	graphics.InitWindow(1200, 800, "Hello from Go!")
	graphics.SetTargetFPS(60)
	defer graphics.CloseWindow()

	frameCount := 0
	for !graphics.WindowShouldClose() {
		frameCount++

		if frameCount%12 == 0 {
			game.calculateNextState()
			game.Swap()
		}

		graphics.BeginDrawing()
		graphics.ClearBackground(color.Black)
		for y, row := range game.state {
			for x := range row {
				if game.At(Point{x, y}) {
					graphics.DrawCell(x, y, cellSize, color.White)
				}
			}
		}
		graphics.EndDrawing()
	}
}
