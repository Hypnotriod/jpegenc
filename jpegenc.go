package jpegenc

/*
#cgo CFLAGS: -g -Wall -O3
#include "c/jpegenc.c"
#include "c/jpegenc.h"
#include "c/jpegenc.inl"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type JpegPixelType int

const (
	JpegPixelTypeGrayscale JpegPixelType = iota
	JpegPixelTypeRGB565    JpegPixelType = iota
	JpegPixelTypeRGB888    JpegPixelType = iota
	JpegPixelTypeARGB8888  JpegPixelType = iota
)

type JpegSubsample int

const (
	JpegSubsample444 JpegSubsample = iota
	JpegSubsample424 JpegSubsample = iota
)

type JpegQualityFactor int

const (
	JpegQualityFactorBest   JpegQualityFactor = iota
	JpegQualityFactorHigh   JpegQualityFactor = iota
	JpegQualityFactorMedium JpegQualityFactor = iota
	JpegQualityFactorLow    JpegQualityFactor = iota
)

type JpegErrorCode int

const (
	JpegSuccess                     JpegErrorCode = iota
	JpegInvalidParameterErrorCode   JpegErrorCode = iota
	JpegEncodeErrorCode             JpegErrorCode = iota
	JpegMemoryErrorCode             JpegErrorCode = iota
	JpegNoBufferErrorCode           JpegErrorCode = iota
	JpegUnsupportedFeatureErrorCode JpegErrorCode = iota
	JpegInvalidFileErrorCode        JpegErrorCode = iota
)

var (
	ErrInvalidParameter   error = errors.New("jpegenc: invalid parameter")
	ErrEncodeError        error = errors.New("jpegenc: encode error")
	ErrMemoryError        error = errors.New("jpegenc: memory error")
	ErrNoBufferError      error = errors.New("jpegenc: no buffer error")
	ErrUnsupportedFeature error = errors.New("jpegenc: unsupported feature")
	ErrInvalidFile        error = errors.New("jpegenc: invalid file")
)

type JpegEncodeParams struct {
	QualityFactor JpegQualityFactor
	PixelType     JpegPixelType
	Subsample     JpegSubsample
}

func Encode(width int, height int, params JpegEncodeParams, pixels []byte, buffer []byte) (bytesEncoded int, err error) {
	errorCode := C.JPEGEncode(
		C.int(width),
		C.int(height),
		C.uchar(params.PixelType),
		C.uchar(params.Subsample),
		C.uchar(params.QualityFactor),
		(*C.uchar)(unsafe.Pointer(&pixels[0])),
		(*C.uchar)(unsafe.Pointer(&buffer[0])),
		C.int(len(buffer)),
		(*C.int)(unsafe.Pointer(&bytesEncoded)),
	)
	switch JpegErrorCode(errorCode) {
	case JpegInvalidParameterErrorCode:
		err = ErrInvalidParameter
	case JpegEncodeErrorCode:
		err = ErrEncodeError
	case JpegMemoryErrorCode:
		err = ErrMemoryError
	case JpegNoBufferErrorCode:
		err = ErrNoBufferError
	case JpegUnsupportedFeatureErrorCode:
		err = ErrUnsupportedFeature
	case JpegInvalidFileErrorCode:
		err = ErrInvalidFile
	}
	return
}
