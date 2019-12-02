package imageprep

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"os"
	"path/filepath"
)

func ImagePrep(name string) {
	const X = 300
	const Y = 149
	var frames []string

	// Get array of paths for frames.
	err := filepath.Walk("resources/frames", func(path string, info os.FileInfo, err error) error {
		frames = append(frames, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(frames); i++ {
		fmt.Printf("%02d\n", i)
	}

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
	dc.DrawStringAnchored(name, 197, 132, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.jpg")
}
