package imageprep

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//ImagePrep accepts name as a string and preps all individual frame for use in final gif.
func ImagePrep(name string) {
	// Size of individual frames.
	const X = 300
	const Y = 149
	var frames []string

	// Get array of paths for frames.
	err := filepath.Walk("resources/frames", func(path string, info os.FileInfo, err error) error {
		frames = append(frames, path)[1:]
		return nil
	})
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v", frames)
	for frame := range frames {
		frame := strings.Split(frames[frame], "/")
		iframe := frame[len(frame)-1]
		fmt.Printf("frame: %s\n", iframe)

		im, err := gg.LoadImage(fmt.Sprintf("resources/frames/%s", iframe))
		if err != nil {
			log.Fatal(err)
		}

		dc := gg.NewContext(X, Y)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
		if err := dc.LoadFontFace("resources/Cocktail.ttf", 20); err != nil {
			panic(err)
		}

		dc.DrawRoundedRectangle(0, 0, X, Y, 0)
		dc.DrawImage(im, 0, 0)
		dc.DrawStringAnchored(strings.ToUpper(name), 199, 133, 0.5, 0.5)
		dc.Clip()
		dc.SavePNG(fmt.Sprintf("resources/outputframes/%s", iframe))
	}
}
