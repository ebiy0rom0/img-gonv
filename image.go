package main

import (
	"flag"
	"fmt"
	"image"
	"math"
	"os"
)

var (
	path         string
	px, py, w, h int
)

func init() {
	flag.StringVar(&path, "path", "", "convert image file path")
	flag.IntVar(&px, "px", math.MaxInt, "reference point on horizontal axis")
	flag.IntVar(&py, "py", math.MaxInt, "reference point on vertical axis")
	flag.IntVar(&w, "w", math.MaxInt, "horizontal distance from px")
	flag.IntVar(&h, "h", math.MaxInt, "vertical distance from py")
}

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, name, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[debug]filename: %s", name)

	return img, nil
}
