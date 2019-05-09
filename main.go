package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

const IMAGESIZE = 2000
const ZOOM = 1
const CENTER = -1 + 0i
const MAXITERATIONS = 1000

type mandelParameters struct {
	output        *image.RGBA
	imagesize     int
	zoom          float64
	center        complex128
	maxiterations int
	segments      int
}

type mandelMaker func(params mandelParameters)

func main() {
	// naive (line-by-line) algorithm
	mandelParams := mandelParameters{imagesize: IMAGESIZE, zoom: ZOOM, center: CENTER, maxiterations: MAXITERATIONS}
	generateImage("naive.png", naiveMandel, mandelParams)

	// concurrent algorithm
	for segmentPow := 1; segmentPow < 6; segmentPow++ {
		segments := int(math.Pow(2, float64(segmentPow)))
		mandelParams = mandelParameters{imagesize: IMAGESIZE, zoom: ZOOM, center: CENTER, maxiterations: MAXITERATIONS,
			segments: segments}
		filename := fmt.Sprintf("concurrent%d.png", segments)
		generateImage(filename, concurrentMandel, mandelParams)
	}
}

func generateImage(fileName string, mandelFunction mandelMaker, mandelParams mandelParameters) {
	outputImage := image.NewRGBA(image.Rect(0, 0, mandelParams.imagesize, mandelParams.imagesize))
	mandelParams.output = outputImage

	mandelFunction(mandelParams)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: improper error handling here
	defer f.Close()
	png.Encode(f, outputImage)
}
