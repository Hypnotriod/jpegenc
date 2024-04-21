#include <string.h>
#include <stdio.h>
#include <stdint.h>
#include "jpegenc.h"

// A simple wrapper function that does just one thing – encodes the raw data into a jpeg
int JPEGEncode(int width, int height, uint8_t pixel_type, uint8_t sub_sample, uint8_t q_factor, bool chroma_swap, uint8_t *pixels, uint8_t *buff, int buff_size, int *bytes_encoded)
{
    int i;
    int code;
    int pixel_size;
    int mcu_count;
    int row_width;
    JPEGIMAGE jpeg_image;
    JPEGENCODE jpeg_encode;

    *bytes_encoded = 0;

    switch (pixel_type)
    {
    case JPEG_PIXEL_GRAYSCALE:
        pixel_size = 1;
        row_width = width * 1;
        break;
    case JPEG_PIXEL_RGB565:
        pixel_size = 2;
        row_width = width * 2;
        break;
    case JPEG_PIXEL_RGB888:
        pixel_size = 3;
        row_width = width * 3;
        break;
    case JPEG_PIXEL_ARGB8888:
        pixel_size = 4;
        row_width = width * 4;
        break;
    default:
        pixel_size = 0;
        row_width = 0;
        return JPEG_UNSUPPORTED_FEATURE;
    }

    memset(&jpeg_image, 0, sizeof(JPEGIMAGE));
    jpeg_image.pOutput = buff;
    jpeg_image.iBufferSize = buff_size;
    jpeg_image.pHighWater = &buff[buff_size - 512];

    code = JPEGEncodeBegin(&jpeg_image,
                           &jpeg_encode,
                           width,
                           height,
                           pixel_type,
                           sub_sample,
                           q_factor,
                           chroma_swap);
    if (code != JPEG_SUCCESS)
        return code;

    mcu_count = ((width + jpeg_encode.cx - 1) / jpeg_encode.cx) * ((height + jpeg_encode.cy - 1) / jpeg_encode.cy);

    for (i = 0; i < mcu_count; i++)
    {
        code = JPEGAddMCU(&jpeg_image,
                          &jpeg_encode,
                          &pixels[(jpeg_encode.x * pixel_size) + (jpeg_encode.y * row_width)],
                          row_width);
        if (code != JPEG_SUCCESS)
            return code;
    }

    JPEGEncodeEnd(&jpeg_image);
    if (code != JPEG_SUCCESS)
        return code;

    *bytes_encoded = jpeg_image.iDataSize;

    return JPEG_SUCCESS;
}
