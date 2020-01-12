package imageprep

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

func getfilepaths(path string) (paths []string) {
	fmt.Println("getting paths.")
	err := filepath.Walk("resources/frames", func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

//ImagePrep accepts name as a string and preps all individual frame for use in final gif.
func ImagePrep(name string) {
	// Size of individual frames.
	const X = 300
	const Y = 149

	// Get array of paths for frames.
	frames := getfilepaths("resources/frames")
	for i, f := range frames[1:] {

		frame := strings.Split(f, "/")[2]
		fmt.Println(i, f, frame)

		im, err := gg.LoadImage(fmt.Sprintf("resources/frames/%s", frame))
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
		switch {
		case i > 1 && i < 13:
			text := fmt.Sprintf("GOD I HATE %s", name)
			dc.DrawStringAnchored(strings.ToUpper(text), 150, 139, 0.5, 0.5)
			dc.Clip()
		}

		dc.SavePNG(fmt.Sprintf("resources/outputframes/%s", frame))
	}
}

//CreateGIF turns images stored in outputframes folder into a .gif
func CreateGIF() {
	fmt.Println("lul")
}
