package main

import (
	"image/color"
	"runtime"

	"github.com/printchard/cgolang/graphics"
)

const width = 100
const height = 100

const screenHeight = 800
const screenWidth = 1200

type Point struct {
	x, y int
}

func (p Point) Add(op Point) Point {
	return Point{p.x + op.x, p.y + op.y}
}

type Game struct {
	state     [][]bool
	nextState [][]bool
	isPaused  bool
	showGrid  bool
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
	return &Game{state: state, nextState: nextState}
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

// SeedGliderGun places the classic Gosper Glider Gun.
// Requires at least a 40x40 grid to run safely without hitting borders immediately.
func (g *Game) SeedGliderGun(origin Point) {
	offsets := []Point{
		// Left square block
		{1, 5}, {2, 5},
		{1, 6}, {2, 6},

		// Left gun shape
		{11, 5}, {11, 6}, {11, 7},
		{12, 4}, {12, 8},
		{13, 3}, {13, 9},
		{14, 3}, {14, 9},
		{15, 6},
		{16, 4}, {16, 8},
		{17, 5}, {17, 6}, {17, 7},
		{18, 6},

		// Right gun shape
		{21, 3}, {21, 4}, {21, 5},
		{22, 3}, {22, 4}, {22, 5},
		{23, 2}, {23, 6},
		{25, 1}, {25, 2}, {25, 6}, {25, 7},

		// Right square block
		{35, 3}, {36, 3},
		{35, 4}, {36, 4},
	}

	for _, offset := range offsets {
		p := origin.Add(offset)
		if p.y >= 0 && p.y < len(g.state) && p.x >= 0 && p.x < len(g.state[0]) {
			g.state[p.y][p.x] = true
		}
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	game := initGame(width, height)

	game.SeedGliderGun(Point{2, 2})

	cellSize := 15

	graphics.InitWindow(screenWidth, screenHeight, "Hello from Go!")
	graphics.SetTargetFPS(60)
	graphics.SetExitKey(graphics.KeyQ)
	defer graphics.CloseWindow()

	frameCount := 0
	for !graphics.WindowShouldClose() {
		frameCount++

		if graphics.IsKeyPressed(graphics.KeySpace) {
			game.isPaused = !game.isPaused
		}
		if graphics.IsKeyPressed(graphics.KeyG) {
			game.showGrid = !game.showGrid
		}

		if frameCount%6 == 0 && !game.isPaused {
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

		if game.showGrid {
			graphics.DrawGrid(width, height, cellSize)
		}

		graphics.EndDrawing()
	}
}
