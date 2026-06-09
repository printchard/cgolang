package main

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/printchard/cgolang/engine"
	"github.com/printchard/cgolang/graphics"
)

const width = 100
const height = 100

const screenHeight = 800
const screenWidth = 1200
const fontSize = 24

const cellSize = 7

func init() {
	runtime.LockOSThread()
}

func main() {
	game := engine.InitGame(width, height)
	game.SeedPattern(engine.Point{X: 2, Y: 2}, engine.GliderGun)

	graphics.InitWindow(screenWidth, screenHeight, "Hello from Go!")
	graphics.SetTargetFPS(60)
	graphics.SetExitKey(graphics.KeyQ)
	defer graphics.CloseWindow()

	frameCount := 0
	for !graphics.WindowShouldClose() {
		frameCount++

		if graphics.IsKeyPressed(graphics.KeySpace) {
			game.IsPaused = !game.IsPaused
		}
		if graphics.IsKeyPressed(graphics.KeyG) {
			game.ShowGrid = !game.ShowGrid
		}
		if game.IsPaused {
			leftDown := graphics.IsMouseButtonDown(graphics.MouseButtonLeft)
			rightDown := graphics.IsMouseButtonDown(graphics.MouseButtonRight)

			if leftDown || rightDown {
				mouseX := graphics.GetMouseX()
				mouseY := graphics.GetMouseY()
				p := engine.Point{X: mouseX / cellSize, Y: mouseY / cellSize}

				if p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height {
					if leftDown {
						game.Set(p, true)
					} else if rightDown {
						game.Set(p, false)
					}
				}
			}
		}

		if frameCount%6 == 0 && !game.IsPaused {
			game.CalculateNextState()
			game.Swap()
		}

		graphics.BeginDrawing()
		graphics.ClearBackground(color.Black)
		for p := range game.Points() {
			if game.AliveAt(p) {
				graphics.DrawCell(p.X, p.Y, cellSize, color.White)
			}
		}

		if game.ShowGrid {
			graphics.DrawGrid(width, height, cellSize)
		}

		graphics.DrawFPS(screenWidth-30, 0)
		generationText := fmt.Sprintf("Generation: %d", game.Generation)
		populationText := fmt.Sprintf("Population: %d", game.AliveCount())

		maxChars := max(len(generationText), len(populationText))
		const estimatedCharWidth = 15
		maxTextLength := (maxChars * estimatedCharWidth)

		graphics.DrawText(generationText, screenWidth-maxTextLength, 20, fontSize, color.White)
		graphics.DrawText(populationText, screenWidth-maxTextLength, 20+fontSize, fontSize, color.White)
		graphics.EndDrawing()
	}
}
