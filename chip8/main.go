package main

import (
	"errors"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"github.com/ejholmes/chip8"
	"github.com/samuel/rgbmatrix"
)

var (
	offColor = color.RGBA{A: 255}
	onColor  = color.RGBA{R: 100, G: 255, B: 100, A: 255}
)

type RGBMatrixDisplay struct {
	m  *rgbmatrix.Matrix
	pb *rgbmatrix.PixelBuffer
}

func NewRGBMatrixDisplay(rows, chained int) (*RGBMatrixDisplay, error) {
	pb, err := rgbmatrix.NewPixelBuffer()
	if err != nil {
		return nil, err
	}
	m, err := rgbmatrix.New(rows, chained)
	if err != nil {
		pb.Destroy()
		return nil, err
	}
	if m.Width() < 64 || m.Height() < 32 {
		pb.Destroy()
		m.Destroy()
		return nil, errors.New("display must at least 64x32")
	}
	return &RGBMatrixDisplay{
		m:  m,
		pb: pb,
	}, nil
}

func (md *RGBMatrixDisplay) Init() error {
	md.m.Clear()
	return nil
}

func (md *RGBMatrixDisplay) Close() {
	md.pb.Destroy()
	md.m.Destroy()
}

func (md *RGBMatrixDisplay) Render(g *chip8.Graphics) error {
	// t := time.Now()
	md.pb.Clear()
	for y := 0; y < 32; y++ {
		for x, c := range g.Pixels[y*64 : (y+1)*64] {
			if c == 0 {
				md.pb.SetPixel(x, y, offColor)
			} else {
				md.pb.SetPixel(x, y, onColor)
			}
		}
	}
	md.m.BlitPixelBuffer(md.pb)
	// dt := time.Since(t)
	// fmt.Printf("%d ms\n", dt.Nanoseconds()/1e6)
	return nil
}

func main() {
	log.Println("Reading rom...")
	rom, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// options := &chip8.Options{
	// }
	cpu, err := chip8.NewCPU(nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer cpu.Close()

	log.Println("Creating display...")
	cpu.Graphics.Display, err = NewRGBMatrixDisplay(32, 2)
	if err != nil {
		log.Fatal(err)
	}
	defer cpu.Graphics.Display.Close()

	log.Println("Loading rom...")
	if _, err := cpu.LoadBytes(rom); err != nil {
		log.Fatal(err)
	}

	log.Println("Running...")
	if err := cpu.Run(); err != nil {
		log.Println(err)
	}
}
