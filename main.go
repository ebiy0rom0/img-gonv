package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Scanln()
}

func run() error {
	fmt.Println(strings.Join([]string{path, strconv.Itoa(px), strconv.Itoa(py), strconv.Itoa(w), strconv.Itoa(h)}, ", "))
	if _, err := LoadImage(""); err != nil {
		fmt.Println(err)
	}
	return nil
}
