package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func LoadImage(target string) (image.Image, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, ext, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[debug]extension: %s\n", ext)

	return img, nil
}
