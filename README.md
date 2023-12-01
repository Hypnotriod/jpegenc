# jpegenc
CGO bindings package for [JPEGENC](https://github.com/bitbank2/JPEGENC) C library. All credits to its authors. 

## installation
```
go get github.com/Hypnotriod/jpegenc
```

## usage
```go
params := jpegenc.EncodeParams{
  QualityFactor: jpegenc.QualityFactorBest,
  PixelType:     jpegenc.PixelTypeRGB565,
  Subsample:     jpegenc.Subsample444,
}
buffer := make([]byte, max(1024, width*height))
// encode raw pixel bytes directly into buffer, no any data allocations/copying by encoder
bytesEncoded, err := jpegenc.Encode(width, height, params, pixelsRGB656[:], buffer)
err = os.WriteFile("file.jpg", buffer[:bytesEncoded], 0644)
```
