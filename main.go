package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	target string
	px, py uint
	w, h   int
)

func init() {
	flag.StringVar(&target, "target", "", "convert image file or directory")
	flag.UintVar(&px, "px", math.MaxUint, "reference point on horizontal axis")
	flag.UintVar(&py, "py", math.MaxUint, "reference point on vertical axis")
	flag.IntVar(&w, "w", math.MaxInt, "horizontal distance from px")
	flag.IntVar(&h, "h", math.MaxInt, "vertical distance from py")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	for len(target) == 0 {
		fmt.Println("target is empty.")
		fmt.Println("Please input convert image file or directory.")
		fmt.Scanln(&target)
	}
	img, err := LoadImage(target)
	if err != nil {
		return err
	}
	rect := img.Bounds()

	for px < uint(rect.Min.X) || px > uint(rect.Max.X) {
		fmt.Println("x-pos exceeds image size.")
		fmt.Printf("Please input x-pos.(Range: %d - %d)\n", rect.Min.X, rect.Max.X)
		fmt.Scanln(&px)
	}

	for py < uint(rect.Min.Y) || py > uint(rect.Max.Y) {
		fmt.Println("y-pos exceeds image size.")
		fmt.Printf("Please input y-pos.(Range: %d - %d)\n", rect.Min.Y, rect.Max.Y)
		fmt.Scanln(&py)
	}

	low, high := -int(px), rect.Max.X-int(px)
	for w < low || w > high {
		fmt.Println("horizontal distance exceeds image size.")
		fmt.Printf("Please input horizontal distance.(Range: %d - %d)\n", low, high)
		fmt.Scanln(&w)
	}

	low, high = -int(py), rect.Max.Y-int(py)
	for h < low || h > high {
		fmt.Println("vertical distance exceeds image size.")
		fmt.Printf("Please input vertical distance.(Range: %d - %d)\n", low, high)
		fmt.Scanln(&h)
	}

	fmt.Println(target, px, py, w, h)
	return nil
}
