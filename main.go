package main

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/printchard/cgolang/graphics"
)

const width = 100
const height = 100

const screenHeight = 800
const screenWidth = 1200
const fontSize = 24

const cellSize = 7

// Classic Gosper Glider Gun.
// Requires at least a 40x40 grid to run safely without hitting borders immediately.
var GliderGun = []Point{
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

var Glider = []Point{
	{1, 0},
	{2, 1},
	{0, 2}, {1, 2}, {2, 2},
}

func init() {
	runtime.LockOSThread()
}

func main() {
	game := InitGame(width, height)
	game.SeedPattern(Point{2, 2}, GliderGun)

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
		for p := range game.Points() {
			if game.AliveAt(p) {
				graphics.DrawCell(p.x, p.y, cellSize, color.White)
			}
		}

		if game.showGrid {
			graphics.DrawGrid(width, height, cellSize)
		}

		graphics.DrawFPS(screenWidth-30, 0)
		generationText := fmt.Sprintf("Generation: %d", game.generation)
		populationText := fmt.Sprintf("Population: %d", game.AliveCount())

		maxChars := max(len(generationText), len(populationText))
		const estimatedCharWidth = 15
		maxTextLength := (maxChars * estimatedCharWidth)

		graphics.DrawText(generationText, screenWidth-maxTextLength, 20, fontSize, color.White)
		graphics.DrawText(populationText, screenWidth-maxTextLength, 20+fontSize, fontSize, color.White)
		graphics.EndDrawing()
	}
}
