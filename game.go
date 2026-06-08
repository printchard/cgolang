package main

type Game struct {
	state      [][]bool
	nextState  [][]bool
	isPaused   bool
	showGrid   bool
	generation int
}

func (g *Game) At(p Point) bool {
	return g.state[p.y][p.x]
}

func (g *Game) SetNext(p Point, alive bool) {
	g.nextState[p.y][p.x] = alive
}

func InitGame(w, h int) *Game {
	state := make([][]bool, h)
	nextState := make([][]bool, h)
	for i := range state {
		state[i] = make([]bool, w)
		nextState[i] = make([]bool, w)
	}
	return &Game{state: state, nextState: nextState}
}

func (g *Game) calculateNextState() {
	g.clearNextState()
	for y := 0; y < len(g.state); y++ {
		for x := 0; x < len(g.state[0]); x++ {
			p := Point{x, y}
			willBeAlive := g.calculateNextCell(p)
			g.SetNext(p, willBeAlive)
		}
	}
}

func (g *Game) calculateNextCell(p Point) bool {
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
		return nCount == 2 || nCount == 3
	} else {
		return nCount == 3
	}
}

func (g *Game) clearNextState() {
	for _, row := range g.nextState {
		clear(row)
	}
}

func (g *Game) Swap() {
	g.state, g.nextState = g.nextState, g.state
	g.generation++
}

func (g *Game) SeedPattern(origin Point, offsets []Point) {
	for _, offset := range offsets {
		p := origin.Add(offset)
		if p.y >= 0 && p.y < len(g.state) && p.x >= 0 && p.x < len(g.state[0]) {
			g.state[p.y][p.x] = true
		}
	}
}

func (g *Game) AliveCount() int {
	count := 0
	for y := range g.state {
		for x := range g.state[y] {
			if g.state[y][x] {
				count++
			}
		}
	}
	return count
}
