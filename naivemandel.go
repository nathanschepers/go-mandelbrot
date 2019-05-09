package main

import (
	"fmt"
	"image/color"
	"time"
)

func naiveMandel(params mandelParameters) {
	imageSize := params.output.Bounds().Dx()

	defer timeTrack(time.Now(), fmt.Sprintf("naive mandelbrot - size: %d, maxiter: %d", imageSize, params.maxiterations))

	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			c := pointToComplex(x, y, imageSize, params.zoom, params.center)
			if inSet, _ := inMandel(c, params.maxiterations); inSet {
				params.output.Set(x, y, color.RGBA{A: 255})
			} else {
				params.output.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}
}
