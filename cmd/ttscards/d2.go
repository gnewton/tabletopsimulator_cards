package main

import (
	"log"

	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image/jpeg"
	"os"
)

func d2() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	w := 640
	h := 1280

	face := truetype.NewFace(font, &truetype.Options{Size: 300})

	dc := gg.NewContext(w, h)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("55", 320, 640, 0.5, 0.5)

	//dc.SavePNG("out.png")

	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var opt jpeg.Options

	opt.Quality = 80

	err = jpeg.Encode(out, dc.Image(), &opt) // put quality to 80%
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
