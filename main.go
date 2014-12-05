package main

import (
	"flag"
)

func main() {
	root := flag.String("path", ".", "path to list")
	depth := flag.Int("depth", 10, "maximal depth")
	size := flag.Int64("size", 0, "min dir size to expand")
	flag.Parse()

}
