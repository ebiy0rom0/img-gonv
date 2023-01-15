package main

import (
	"flag"
	"fmt"

	"github.com/ebiy0rom0/img-gonv/convert"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	if err := convert.C.ReadImage(); err != nil {
		return err
	}
	if err := convert.C.Convert(); err != nil {
		return err
	}
	return nil
}
