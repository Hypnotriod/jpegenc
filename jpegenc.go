package jpegenc

/*
#cgo CFLAGS: -g -Wall -O3 -march=native
#include "c/jpegenc.c"
#include "c/jpegenc.h"
#include "c/jpegenc.inl"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type PixelType int

const (
	PixelTypeGrayscale PixelType = iota
	PixelTypeRGB565    PixelType = iota
	PixelTypeRGB888    PixelType = iota
	PixelTypeARGB8888  PixelType = iota
)

type Subsample int

const (
	Subsample444 Subsample = iota
	Subsample424 Subsample = iota
)

type QualityFactor int

const (
	QualityFactorBest   QualityFactor = iota
	QualityFactorHigh   QualityFactor = iota
	QualityFactorMedium QualityFactor = iota
	QualityFactorLow    QualityFactor = iota
)

type ErrorCode int

const (
	Success                     ErrorCode = iota
	InvalidParameterErrorCode   ErrorCode = iota
	EncodeErrorCode             ErrorCode = iota
	MemoryErrorCode             ErrorCode = iota
	NoBufferErrorCode           ErrorCode = iota
	UnsupportedFeatureErrorCode ErrorCode = iota
	InvalidFileErrorCode        ErrorCode = iota
)

var (
	ErrInvalidParameter   error = errors.New("jpegenc: invalid parameter")
	ErrEncodeError        error = errors.New("jpegenc: encode error")
	ErrMemoryError        error = errors.New("jpegenc: memory error")
	ErrNoBufferError      error = errors.New("jpegenc: no buffer error")
	ErrUnsupportedFeature error = errors.New("jpegenc: unsupported feature")
	ErrInvalidFile        error = errors.New("jpegenc: invalid file")
)

type EncodeParams struct {
	QualityFactor QualityFactor
	PixelType     PixelType
	Subsample     Subsample
}

func Encode(width int, height int, params EncodeParams, pixels []byte, buffer []byte) (bytesEncoded int, err error) {
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
	switch ErrorCode(errorCode) {
	case InvalidParameterErrorCode:
		err = ErrInvalidParameter
	case EncodeErrorCode:
		err = ErrEncodeError
	case MemoryErrorCode:
		err = ErrMemoryError
	case NoBufferErrorCode:
		err = ErrNoBufferError
	case UnsupportedFeatureErrorCode:
		err = ErrUnsupportedFeature
	case InvalidFileErrorCode:
		err = ErrInvalidFile
	}
	return
}
