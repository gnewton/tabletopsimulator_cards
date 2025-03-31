package main

import (
	//"errors"
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"log"
	"path/filepath"
	//"github.com/llgcode/draw2d/draw2dkit"
	//"github.com/llgcode/draw2d/samples"
	//"github.com/llgcode/draw2d/samples/gopher2"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"os"
	"strconv"
)

func createTestImages(backFlag string, imageDirectoryFlag string, w, h int) error {
	// Test directory exist?
	if _, err := os.Stat(imageDirectoryFlag); os.IsNotExist(err) {
		err = os.MkdirAll(imageDirectoryFlag, 0755)
		if err != nil {
			return err
		}
	}

	// Delete all files from test fir
	err := DeleteDirectoryFiles(imageDirectoryFlag)
	if err != nil {
		return err
	}

	numCards := DEFAULT_NUM_ROWS_CARDS*DEFAULT_NUM_COLUMNS_CARDS - 1
	for i := 0; i < numCards; i++ {
		var cardNumber string
		if i < 10 {
			cardNumber = "0"
		}
		cardNumber = cardNumber + strconv.Itoa(i)
		makeCardImage(imageDirectoryFlag+"/"+cardNumber+".png", cardNumber, w, h)
	}

	makeBackImage(imageDirectoryFlag+"/"+backFlag, w, h)

	return nil
}

var fillToggle = true

func makeCardImage(filename string, cardNumber string, w, h int) {

	verbose(filename)

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	// if fillToggle {
	// 	gc.SetFillColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
	// 	fillToggle = false
	// } else {
	// 	fillToggle = true
	// 	gc.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	// }

	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	//gc.SetLineWidth(5)
	gc.SetFontSize(float64(w / 2))
	gc.SetFontData(draw2d.FontData{
		Name: "gomono",
	})

	gc.Save()
	// if fillToggle {
	// 	gc.SetFillColor(color.RGBA{0x00, 0x00, 0xf0, 0xff})
	// 	fillToggle = false
	// } else {
	// 	gc.SetFillColor(color.RGBA{0xf0, 0x00, 0x00, 0xff})
	// 	fillToggle = true
	// }

	gc.MoveTo(0, 0)
	gc.LineTo(float64(w), 0)
	gc.LineTo(float64(w), float64(h))
	gc.LineTo(0, float64(h))
	gc.Close()

	//gc.Fill()
	gc.Save()
	gc.SetFillColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
	gc.Fill()
	gc.Restore()

	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.Stroke()
	gc.Restore()

	gc.Save()
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	gc.MoveTo(0, 0)
	gc.Translate(float64(w)/10, float64(h)/1.5)

	gc.FillString(cardNumber)
	gc.Close()
	gc.FillStroke()

	gc.Restore()

	// Save to file
	err := draw2dimg.SaveToPngFile(filename, dest)
	if err != nil {
		log.Fatal(err)
	}

}

// // From https://github.com/llgcode/draw2d/issues/127
type MyFontCache map[string]*truetype.Font

func (fc MyFontCache) Store(fd draw2d.FontData, font *truetype.Font) {
	fc[fd.Name] = font
}

func (fc MyFontCache) Load(fd draw2d.FontData) (*truetype.Font, error) {
	font, stored := fc[fd.Name]
	if !stored {
		return nil, fmt.Errorf("Font %s is not stored in font cache.", fd.Name)
	}
	return font, nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fontCache := MyFontCache{}

	TTFs := map[string]([]byte){
		"goregular": goregular.TTF,
		"gobold":    gobold.TTF,
		"goitalic":  goitalic.TTF,
		"gomono":    gomono.TTF,
	}

	for fontName, TTF := range TTFs {
		font, err := truetype.Parse(TTF)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		fontCache.Store(draw2d.FontData{Name: fontName}, font)
	}

	draw2d.SetFontCache(fontCache)
}

func DeleteDirectoryFiles(dir string) error {
	files, err := filepath.Glob(dir + "/*")
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

func makeBackImage(filename string, w, h int) {
	verbose(filename)

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	//gc.SetLineWidth(5)
	gc.SetFontSize(float64(w / 4))
	gc.SetFontData(draw2d.FontData{
		Name: "gomono",
	})

	gc.Save()
	gc.MoveTo(0, 0)
	gc.Translate(float64(w)/10, float64(h)/1.75)

	gc.FillString("back")
	gc.Close()
	gc.FillStroke()

	gc.Restore()

	// Save to file
	err := draw2dimg.SaveToPngFile(filename, dest)
	if err != nil {
		log.Fatal(err)
	}

}
