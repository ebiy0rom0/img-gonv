package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	target       string
	px, py, w, h uint
)

func init() {
	flag.StringVar(&target, "target", "", "convert image file or directory")
	flag.UintVar(&px, "px", math.MaxUint, "reference point on horizontal axis")
	flag.UintVar(&py, "py", math.MaxUint, "reference point on vertical axis")
	flag.UintVar(&w, "w", math.MaxUint, "horizontal distance from px")
	flag.UintVar(&h, "h", math.MaxUint, "vertical distance from py")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Scanln()
}

func run() error {
	for len(target) == 0 {
		fmt.Println("target is empty.")
		fmt.Println("Please input convert image file or directory.")
		fmt.Scan(&target)
	}
	img, err := LoadImage(target)
	if err != nil {
		return err
	}
	rect := img.Bounds()

	fmt.Println(rect)
	return nil
}
