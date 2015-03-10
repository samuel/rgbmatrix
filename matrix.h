#ifndef _MATRIX_H_
#define _MATRIX_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "buffer.h"

// typedef void* CRGBMatrix;
typedef struct {
	void* gpio;
	void* matrix;
} CRGBMatrix;

CRGBMatrix NewRGBMatrix(int rows, int chainedDisplays, int *error);
void FreeRGBMatrix(CRGBMatrix mat);

// // Set PWM bits used for output. Default is 11, but if you only deal with
// // simple comic-colors, 1 might be sufficient. Lower require less CPU.
// // Returns boolean to signify if value was within range.
// bool SetPWMBits(uint8_t value);
// uint8_t pwmbits();

// // Map brightness of output linearly to input with CIE1931 profile.
// void set_luminance_correct(bool on);
// bool luminance_correct() const;

int RGBMatrixWidth(CRGBMatrix mat);
int RGBMatrixHeight(CRGBMatrix mat);
void RGBMatrixSetPixel(CRGBMatrix mat, int x, int y, uint8_t red, uint8_t green, uint8_t blue);
void RGBMatrixClear(CRGBMatrix mat);
void RGBMatrixFill(CRGBMatrix mat, uint8_t red, uint8_t green, uint8_t blue);

void RGBMatrixBlitPixelBuffer(CRGBMatrix mat, PixelBuffer pb);

#ifdef __cplusplus
}
#endif

#endif
