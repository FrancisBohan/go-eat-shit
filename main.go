package main

import img "github.com/francisbohan/go-eat-shit/imageprep"

func main() {
	name := "frank"
	img.ImagePrep(name)
	img.CreateGIF(name)
}
