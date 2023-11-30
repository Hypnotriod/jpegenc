# jpegenc
CGO bindings package for [JPEGENC](https://github.com/bitbank2/JPEGENC) C library. All credits to its authors. 

## installation
```
go get github.com/Hypnotriod/jpegenc
```

## usage
```go
params := jpegenc.JpegEncodeParams{
  QualityFactor: jpegenc.JpegQualityFactorBest,
  PixelType:     jpegenc.JpegPixelTypeRGB565,
  Subsample:     jpegenc.JpegSubsample444,
}
// data will be stored directly in buffer slice, no any allocations by encoder
buffer := make([]byte, max(1024, width*height))
// encode raw pixel bytes, no any data copying by encoder
bytesEncoded, err := jpegenc.Encode(width, height, params, pixelsRGB656[:], buffer)
err = os.WriteFile("file.jpg", buffer[:bytesEncoded], 0644)
```
