package convert

import (
	"flag"
	"fmt"
	"image"
	"io/fs"
	"math"
	"os"
	"path/filepath"

	"image/draw"
	"image/jpeg"
	_ "image/png"

	"github.com/ebiy0rom0/img-gonv/config"
)

type imageInfo struct {
	name string
	img  image.Image
}

type converter struct {
	img    []imageInfo
	target string
	isDir  bool
	px, py uint
	w, h   int
}

var C = &converter{}

func init() {
	flag.StringVar(&C.target, "target", "", "convert image file or directory")
	flag.StringVar(&C.target, "t", "", "shorthand of target")
	flag.UintVar(&C.px, "pos-x", math.MaxUint, "reference point on horizontal axis")
	flag.UintVar(&C.px, "px", math.MaxUint, "shorthand of pos-x")
	flag.UintVar(&C.py, "pos-y", math.MaxUint, "reference point on vertical axis")
	flag.UintVar(&C.py, "py", math.MaxUint, "shorthand of pos-y")
	flag.IntVar(&C.w, "width", math.MaxInt, "horizontal distance from px")
	flag.IntVar(&C.w, "w", math.MaxInt, "shorthand of width")
	flag.IntVar(&C.h, "height", math.MaxInt, "vertical distance from py")
	flag.IntVar(&C.h, "h", math.MaxInt, "shotrhand of height")
}

func (c *converter) ReadImage() error {
	for len(c.target) == 0 {
		fmt.Println("target is empty.")
		fmt.Println("Please input convert image file or directory.")
		fmt.Scanln(&c.target)

		if f, err := os.Stat(c.target); os.IsNotExist(err) {
			fmt.Println(err)
			c.target = ""
		} else {
			c.isDir = f.IsDir()
		}
	}
	return c.open()
}

func (c *converter) open() error {
	if c.isDir {
		return c.openDir()
	}
	return c.openFile(c.target)
}

func (c *converter) openFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, ext, err := image.Decode(file)
	if err != nil {
		return err
	}
	fmt.Printf("[debug]extension: %s\n", ext)
	_, filename := filepath.Split(path)
	c.img = append(c.img, imageInfo{name: filename, img: img})

	return nil
}

func (c *converter) openDir() error {
	if err := filepath.Walk(c.target, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			return nil
		}

		if err := c.openFile(path); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *converter) Convert() error {
	for _, info := range c.img {
		if err := c.convert(info); err != nil {
			return err
		}
	}
	return nil
}

func (c *converter) convert(info imageInfo) error {
	rect := info.img.Bounds()

	for c.px < uint(rect.Min.X) || c.px > uint(rect.Max.X) {
		fmt.Println("x-pos exceeds image size.")
		fmt.Printf("Please input x-pos.(Range: %d ~ %d)\n", rect.Min.X, rect.Max.X)
		fmt.Scanln(&c.px)
	}

	for c.py < uint(rect.Min.Y) || c.py > uint(rect.Max.Y) {
		fmt.Println("y-pos exceeds image size.")
		fmt.Printf("Please input y-pos.(Range: %d ~ %d)\n", rect.Min.Y, rect.Max.Y)
		fmt.Scanln(&c.py)
	}

	min, max := -int(c.px), rect.Max.X-int(c.px)
	for c.w < min || c.w > max {
		fmt.Println("horizontal distance exceeds image size.")
		fmt.Printf("Please input horizontal distance.(Range: %d ~ %d)\n", min, max)
		fmt.Scanln(&c.w)
	}

	min, max = -int(c.py), rect.Max.Y-int(c.py)
	for c.h < min || c.h > max {
		fmt.Println("vertical distance exceeds image size.")
		fmt.Printf("Please input vertical distance.(Range: %d ~ %d)\n", min, max)
		fmt.Scanln(&c.h)
	}

	minX, maxX := c.orderAsc(int(c.px), int(c.px)+c.w)
	minY, maxY := c.orderAsc(int(c.py), int(c.py)+c.h)

	trimRect := image.Rectangle{
		Min: image.Point{X: minX, Y: minY},
		Max: image.Point{X: maxX, Y: maxY},
	}

	dst := image.NewRGBA(trimRect)
	draw.Draw(dst, rect, info.img, image.Point{}, draw.Over)

	convPath := filepath.Join(config.OutputPath, info.name)
	file, err := os.Create(convPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := jpeg.Encode(file, dst, &jpeg.Options{Quality: 100}); err != nil {
		return err
	}

	return nil
}

func (c *converter) orderAsc(n1, n2 int) (int, int) {
	if n1 <= n2 {
		return n1, n2
	}
	return n2, n1
}
