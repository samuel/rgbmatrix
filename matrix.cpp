// #include <vector>
#include "led-matrix.h"
#include "matrix.h"

using rgb_matrix::GPIO;
using rgb_matrix::RGBMatrix;

CRGBMatrix NewRGBMatrix(int rows, int chainedDisplays, int *error) {
	*error = 0;
	CRGBMatrix mat = {NULL, NULL};
	GPIO *gpio = new GPIO();
	if (!gpio->Init()) {
		delete gpio;
		*error = 1;
		return mat;
	}
	mat.gpio = gpio;
	mat.matrix = new RGBMatrix(gpio, rows, chainedDisplays);
	return mat;
}

void FreeRGBMatrix(CRGBMatrix mat) {
	if (mat.gpio != NULL) {
		delete (GPIO*)mat.gpio;
	}
	if (mat.matrix != NULL) {
		delete (RGBMatrix*)mat.matrix;
	}
}

int RGBMatrixWidth(CRGBMatrix mat) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	return m->width();
}

int RGBMatrixHeight(CRGBMatrix mat) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	return m->height();
}

void RGBMatrixSetPixel(CRGBMatrix mat, int x, int y, uint8_t red, uint8_t green, uint8_t blue) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	m->SetPixel(x, y, red, green, blue);
}

void RGBMatrixClear(CRGBMatrix mat) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	m->Clear();
}

void RGBMatrixFill(CRGBMatrix mat, uint8_t red, uint8_t green, uint8_t blue) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	m->Fill(red, green, blue);
}

void RGBMatrixBlitPixelBuffer(CRGBMatrix mat, PixelBuffer pb) {
	RGBMatrix *m = (RGBMatrix*)mat.matrix;
	for(size_t i = 0; i < pb.size; i++) {
		Pixel p = pb.buf[i];
		m->SetPixel(p.x, p.y, p.red, p.green, p.blue);
	}
}
