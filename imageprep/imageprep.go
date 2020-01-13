package imageprep

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

//GoEatShit ...
func GoEatShit(name string) {
	ImagePrep(name)
	CreateGIF(name)
}

func getfilepaths(folderpath string) (paths []string) {
	err := filepath.Walk(folderpath, func(path string, info os.FileInfo, err error) error {
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
		case i > 1 && i < 15:
			text := fmt.Sprintf("GOD I HATE %s", name)
			dc.DrawStringAnchored(strings.ToUpper(text), 150, 139, 0.5, 0.5)
		case i > 18 && i < 27:
			text := fmt.Sprintf("HEY %s", name)
			dc.DrawStringAnchored(strings.ToUpper(text), 150, 139, 0.5, 0.5)
		case i > 48:
			text := fmt.Sprintf("YOU WIN THIS ROUND, %s", name)
			dc.DrawStringAnchored(strings.ToUpper(text), 150, 139, 0.5, 0.5)
		}
		dc.Clip()

		x := dc.Image()
		out, err := os.Create(fmt.Sprintf("resources/outputframes/%s.gif", frame))
		if err != nil {
			panic(err)
		}

		var opt gif.Options
		opt.NumColors = 255
		err = gif.Encode(out, x, &opt)
		if err != nil {
			panic(err)
		}
	}
}

//CreateGIF turns images stored in outputframes folder into a .gif
func CreateGIF(name string) {
	frames := getfilepaths("resources/outputframes")
	outGif := &gif.GIF{}
	for _, name := range frames[1:] {
		delay, err := strconv.Atoi(strings.Replace((strings.Split(name, ".")[1]), "s", "", 1))
		if err != nil {
			panic(err)
		}
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		img, err := gif.Decode(f)
		if err != nil {
			panic(err)
		}
		f.Close()

		outGif.Image = append(outGif.Image, img.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, delay)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s.gif", name), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
