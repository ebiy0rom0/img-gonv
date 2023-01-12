package main

import (
	"fmt"
	"image"
	"os"
)

func LoadImage(target string) (image.Image, error) {
	file, err := os.Open(target)
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
