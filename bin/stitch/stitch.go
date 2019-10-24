// Copyright (c) 2019 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var inputPath = filepath.Join(".", "..", "..", "output")

func main() {
	log.Printf("Starting to read tile information at \"%v\"", inputPath)
	tiles, err := loadImages(inputPath, 2)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Got %v tiles", len(tiles))

	totalBounds := image.Rectangle{}
	for i, tile := range tiles {
		if i == 0 {
			totalBounds = tile.Bounds()
		} else {
			totalBounds = totalBounds.Union(tile.Bounds())
		}
	}
	log.Printf("Total size of the possible output space is %v", totalBounds)

	/*profFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Panicf("could not create CPU profile: %v", err)
	}
	defer profFile.Close()
	if err := pprof.StartCPUProfile(profFile); err != nil {
		log.Panicf("could not start CPU profile: %v", err)
	}
	defer pprof.StopCPUProfile()*/

	// TODO: Flags / Program arguments

	outputRect := image.Rect(-1900*2, -1900*2, 1900*2, 1900*2)

	log.Printf("Creating output image with a size of %v", outputRect.Size())
	outputImage := image.NewRGBA(outputRect)

	log.Printf("Stitching %v tiles into an image at %v", len(tiles), outputImage.Bounds())
	tp := make(tilePairs)
	if err := tp.StitchGrid(tiles, outputImage, 512); err != nil {
		log.Panic(err)
	}

	log.Printf("Creating output file \"%v\"", "output.png")
	f, err := os.Create("output.png")
	if err != nil {
		log.Panic(err)
	}

	if err := png.Encode(f, outputImage); err != nil {
		f.Close()
		log.Panic(err)
	}

	if err := f.Close(); err != nil {
		log.Panic(err)
	}
	log.Printf("Created output file \"%v\"", "output.png")

}
