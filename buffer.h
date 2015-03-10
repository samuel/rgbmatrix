#ifndef _BUFFER_H_
#define _BUFFER_H_

#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>

typedef unsigned char uint8_t;

typedef struct {
	int x;
	int y;
	uint8_t red;
	uint8_t green;
	uint8_t blue;
} Pixel;

typedef struct {
	size_t capacity;
	size_t size;
	Pixel *buf;
} PixelBuffer;

#ifdef __cplusplus
}
#endif

#endif
