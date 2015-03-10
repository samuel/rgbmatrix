package rgbmatrix

// #include <stdlib.h>
// #include "buffer.h"
import "C"

import (
	"errors"
	"image/color"
	"unsafe"
)

var pixelSize C.size_t

func init() {
	var p C.Pixel
	pixelSize = C.size_t(unsafe.Sizeof(p))
}

type PixelBuffer struct {
	cpb C.PixelBuffer
	buf []C.Pixel
}

func NewPixelBuffer() (*PixelBuffer, error) {
	initialCapacity := 64 * 64 // Likely max we would need
	cbuf := (*C.Pixel)(C.malloc(C.size_t(initialCapacity) * pixelSize))
	if cbuf == nil {
		return nil, errors.New("failed to allocate C buffer")
	}
	var cpb C.PixelBuffer
	cpb.capacity = C.size_t(initialCapacity)
	cpb.size = C.size_t(0)
	cpb.buf = cbuf
	return &PixelBuffer{
		cpb: cpb,
		buf: (*[1 << 20]C.Pixel)(unsafe.Pointer(cpb.buf))[:cpb.capacity],
	}, nil
}

func (pb *PixelBuffer) Destroy() {
	if pb.cpb.buf != nil {
		C.free(unsafe.Pointer(pb.cpb.buf))
		pb.cpb.buf = nil
		pb.cpb.size = 0
		pb.cpb.capacity = 0
	}
}

func (pb *PixelBuffer) SetPixel(x, y int, c color.RGBA) {
	if pb.cpb.size == pb.cpb.capacity {
		pb.cpb.capacity *= 2
		pb.cpb.buf = (*C.Pixel)(C.realloc(unsafe.Pointer(pb.cpb.buf), pb.cpb.capacity*pixelSize))
		if pb.cpb.buf == nil {
			panic("OOM growing pixel buffer")
		}
		pb.buf = (*[1 << 20]C.Pixel)(unsafe.Pointer(pb.cpb.buf))[:pb.cpb.capacity]
	}
	var p C.Pixel
	p.x = C.int(x)
	p.y = C.int(y)
	p.red = C.uint8_t(c.R)
	p.green = C.uint8_t(c.G)
	p.blue = C.uint8_t(c.B)
	pb.buf[pb.cpb.size] = p
	pb.cpb.size++
}

func (pb *PixelBuffer) Clear() {
	pb.cpb.size = 0
}
