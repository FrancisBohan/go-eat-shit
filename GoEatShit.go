package main

import (
	"flag"

	img "github.com/francisbohan/go-eat-shit/imageprep"
)

func main() {
	name := flag.String("name", "Frank", "")
	flag.Parse()
	img.GoEatShit(*name)
}
