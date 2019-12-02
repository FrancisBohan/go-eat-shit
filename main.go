package main

import (
	"github.com/fogleman/gg"
	"log"
)

func main() {
	const X = 300
	const Y = 149
	im, err := gg.LoadImage("resources/frames/frame_02_delay-0.06s.jpg")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("resources/Cocktail.ttf", 18); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, X, Y, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("FRANK", 197, 132, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.png")
}
