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
	fmt.Println("d2")
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	w := 600
	h := 800

	face := truetype.NewFace(font, &truetype.Options{Size: (float64(w) / 1.8)})

	dc := gg.NewContext(w, h)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("55", float64(w/2), float64(h/2), 0.5, 0.5)

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
