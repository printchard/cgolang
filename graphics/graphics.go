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
