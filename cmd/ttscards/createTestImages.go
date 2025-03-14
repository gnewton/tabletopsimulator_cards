package main

import (
	//"errors"
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"log"
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

	if _, err := os.Stat(imageDirectoryFlag); os.IsNotExist(err) {
		err = os.MkdirAll(imageDirectoryFlag, 0755)
		if err != nil {
			return err
		}
	}

	for i := 0; i < 78; i++ {
		var cardNumber string
		if i < 10 {
			cardNumber = "0"
		}
		cardNumber = cardNumber + strconv.Itoa(i)
		makeCardImage(imageDirectoryFlag+"/"+cardNumber+".png", cardNumber, w, h)
	}

	return nil
}

func makeCardImage(filename string, cardNumber string, w, h int) {

	verbose(filename)

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	//gc.SetLineWidth(5)
	gc.SetFontSize(float64(w / 2))
	gc.SetFontData(draw2d.FontData{
		Name: "gomono",
	})

	// Draw a closed shape
	// gc.MoveTo(100, 500) // should always be called first for a new path
	// gc.LineTo(100, -50)
	// gc.QuadCurveTo(100, 10, 10, 10)
	// gc.Close()
	// gc.FillStroke()

	gc.Save()
	gc.MoveTo(0, 0)
	gc.Translate(float64(w)/10, float64(h)/1.75)

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
			panic(err)
		}
		fontCache.Store(draw2d.FontData{Name: fontName}, font)
	}

	draw2d.SetFontCache(fontCache)
}
