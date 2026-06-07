package graphics

/*
#cgo CFLAGS: -I/opt/homebrew/opt/raylib/include
#cgo LDFLAGS: -L/opt/homebrew/opt/raylib/lib -lraylib
#include <raylib.h>
#include <stdlib.h>
*/
import "C"

import (
	"image/color"
	"unsafe"
)

type Key int

const (
	KeyQ Key = iota
	KeySpace
	KeyG
)

var Grey color.Color = color.RGBA{128, 128, 128, 255}

func InitWindow(width, height int, title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.InitWindow(C.int(width), C.int(height), t)
}

func WindowShouldClose() bool {
	return bool(C.WindowShouldClose())
}

func BeginDrawing() {
	C.BeginDrawing()
}

func EndDrawing() {
	C.EndDrawing()
}

func toCColor(c color.Color) C.Color {
	r, g, b, a := c.RGBA()

	return C.Color{
		r: C.uchar(r >> 8),
		g: C.uchar(g >> 8),
		b: C.uchar(b >> 8),
		a: C.uchar(a >> 8),
	}
}

func ClearBackground(c color.Color) {
	C.ClearBackground(toCColor(c))
}

func CloseWindow() {
	C.CloseWindow()
}

func SetTargetFPS(fps int) {
	C.SetTargetFPS(C.int(fps))
}

func DrawCell(x, y, cellSize int, c color.Color) {
	C.DrawRectangle(C.int(x*cellSize), C.int(y*cellSize), C.int(cellSize), C.int(cellSize), toCColor(c))
}

func DrawGrid(gridWidth, gridHeight, cellSize int) {
	pixelWidth := gridWidth * cellSize
	pixelHeight := gridHeight * cellSize

	for x := 0; x <= pixelWidth; x += cellSize {
		C.DrawLine(
			C.int(x),
			0,
			C.int(x),
			C.int(pixelHeight),
			toCColor(Grey),
		)
	}

	for y := 0; y <= pixelHeight; y += cellSize {
		C.DrawLine(
			0,
			C.int(y),
			C.int(pixelWidth),
			C.int(y),
			toCColor(Grey),
		)
	}
}

func SetExitKey(k Key) {
	switch k {
	case KeyQ:
		C.SetExitKey(C.KEY_Q)
	}
}

func IsKeyPressed(k Key) bool {
	switch k {
	case KeySpace:
		return bool(C.IsKeyPressed(C.KEY_SPACE))
	case KeyG:
		return bool(C.IsKeyPressed(C.KEY_G))
	default:
		return false
	}
}

func DrawFPS(x, y int) {
	C.DrawFPS(C.int(x), C.int(y))
}

func DrawText(text string, x, y, fontSize int, c color.Color) {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))

	C.DrawText(t, C.int(x), C.int(y), C.int(fontSize), toCColor(c))
}

func MeasureText(text string, fontSize int) int {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))

	return int(C.MeasureText(t, C.int(fontSize)))
}
