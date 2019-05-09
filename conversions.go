package main

func pointToComplex(x int, y int, imageSize int, zoom float64, center complex128) complex128 {
	windowSize := 1 / zoom
	scalingFactor := windowSize / float64(imageSize)
	centeredX := float64(x) - float64(imageSize)/2
	centeredY := float64(y) - float64(imageSize)/2
	return complex(centeredX*scalingFactor, centeredY*scalingFactor) + center
}
