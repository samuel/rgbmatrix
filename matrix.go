package rgbmatrix

// #cgo LDFLAGS: -lrgbmatrix
//
// #include "matrix.h"
import "C"

import (
	"errors"
	"image/color"
)

type Matrix struct {
	cm            C.CRGBMatrix
	width, height int
}

func New(rows, chainedDisplays int) (*Matrix, error) {
	var e C.int
	cm := C.NewRGBMatrix(32, 2, &e)
	switch e {
	case 0:
	case 1:
		return nil, errors.New("gpio init failed")
	default:
		return nil, errors.New("unknown error creating matrix")
	}
	if e != 0 {
	}
	width := C.RGBMatrixWidth(cm)
	height := C.RGBMatrixHeight(cm)
	mat := &Matrix{
		cm:     cm,
		width:  int(width),
		height: int(height),
	}
	return mat, nil
}

func (m *Matrix) Destroy() {
	if m.cm.gpio != nil {
		C.FreeRGBMatrix(m.cm)
		m.cm.gpio = nil
	}
}

func (m *Matrix) Width() int {
	return m.width
}

func (m *Matrix) Height() int {
	return m.height
}

func (m *Matrix) SetPixel(x, y int, c color.RGBA) {
	C.RGBMatrixSetPixel(m.cm, C.int(x), C.int(y), C.uint8_t(c.R), C.uint8_t(c.G), C.uint8_t(c.B))
}

func (m *Matrix) Clear() {
	C.RGBMatrixClear(m.cm)
}

func (m *Matrix) Fill(c color.RGBA) {
	C.RGBMatrixFill(m.cm, C.uint8_t(c.R), C.uint8_t(c.G), C.uint8_t(c.B))
}

func (m *Matrix) BlitPixelBuffer(pb *PixelBuffer) {
	C.RGBMatrixBlitPixelBuffer(m.cm, pb.cpb)
}

// func main() {
// 	mat, err := New(32, 2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mat.Destroy()

// 	buf, err := NewPixelBuffer()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer buf.Destroy()

// 	mat.Clear()
// 	mat.SetPixel(0, 0, color.RGBA{R: 100, G: 255, B: 100})
// 	mat.SetPixel(1, 1, color.RGBA{R: 100, G: 255, B: 100})

// 	for i := 0; i < 100; i++ {
// 		buf.SetPixel(
// 			rand.Intn(mat.Width()),
// 			rand.Intn(mat.Height()),
// 			color.RGBA{
// 				R: uint8(rand.Intn(256)),
// 				G: uint8(rand.Intn(256)),
// 				B: uint8(rand.Intn(256)),
// 			})
// 	}
// 	mat.BlitPixelBuffer(buf)

// 	time.Sleep(time.Second * 1)
// 	mat.Fill(color.RGBA{R: 255, G: 0, B: 0})
// 	time.Sleep(time.Second * 1)
// 	mat.Clear()
// }
