package main

import (
	"fmt"
	"image/color"
	"sync"
	"time"
)

func concurrentMandel(params mandelParameters) {
	defer timeTrack(time.Now(), fmt.Sprintf("concurrent mandelbrot - size: %d, maxiter: %d, segments: %d",
		params.imagesize, params.maxiterations, params.segments))

	var wg sync.WaitGroup

	for segment := 0; segment < params.segments; segment++ {
		wg.Add(1)
		go processSegment(segment, params, &wg)
	}

	wg.Wait()
}

func processSegment(segment int, params mandelParameters, wg *sync.WaitGroup) {
	segmentStart := segment * params.imagesize / params.segments
	segmentEnd := (segment + 1) * params.imagesize / params.segments

	for y := segmentStart; y < segmentEnd; y++ {
		for x := 0; x < params.imagesize; x++ {
			c := pointToComplex(x, y, params.imagesize, params.zoom, params.center)
			if inSet, _ := inMandel(c, params.maxiterations); inSet {
				params.output.Set(x, y, color.RGBA{A: 255})
			} else {
				params.output.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}
	wg.Done()
}
