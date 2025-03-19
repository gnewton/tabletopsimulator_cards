package main

import (
	"errors"
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	//"github.com/llgcode/draw2d/draw2dimg/Matrix"
	"image"
	"image/draw"
	"io/fs"
	"log"
	"os"
)

func makeCardsPage(args *Args) error {
	if err := verify(args); err != nil {
		return err
	}

	// Source card images
	if _, err := os.Stat(*args.imageDirectoryFlag); os.IsNotExist(err) {
		return errors.New("Image files directory (source) " + *args.imageDirectoryFlag + " does not exist.")
	}

	root := os.DirFS(*args.imageDirectoryFlag)
	cardFiles, err := fs.ReadDir(root, ".")

	if err != nil {
		log.Fatal(err)
	}

	w, h, err := getCardDimensions(*args.imageDirectoryFlag + cardFiles[0].Name())
	if err != nil {
		return err
	}

	if w**args.numColumnsOfCards > MAX_IMAGE_WIDTH {
		return fmt.Errorf("Too many cards for width: %d columns x %d width = %d which is larger than max allowable width: %d", *args.numColumnsOfCards, w, *args.numColumnsOfCards*w, MAX_IMAGE_WIDTH)
	}

	if w**args.numRowsOfCards > MAX_IMAGE_HEIGHT {
		return fmt.Errorf("Too many cards for height: %d rows x %d height = %d which is larger than max allowable height: %d", *args.numRowsOfCards, h, *args.numRowsOfCards*h, MAX_IMAGE_HEIGHT)
	}

	if err := allCardsSameDimension(w, h, *args.imageDirectoryFlag, cardFiles); err != nil {
		return err
	}

	dest := image.NewRGBA(image.Rect(0, 0, w**args.numColumnsOfCards-1, h**args.numRowsOfCards-1))

	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.MoveTo(0, 0)
	gc.LineTo(float64(w**args.numColumnsOfCards-1), 0)
	gc.LineTo(float64(w**args.numColumnsOfCards-1), float64(h**args.numRowsOfCards-1))
	gc.LineTo(0, float64(h**args.numRowsOfCards-1))
	gc.Close()
	gc.Fill()

	icard := 0
	for j := 0; j < *args.numRowsOfCards; j++ {
		for i := 0; i < *args.numColumnsOfCards; i++ {
			img, err := draw2dimg.LoadFromPngFile(*args.imageDirectoryFlag + cardFiles[icard].Name())
			if err != nil {
				return err
			}
			draw2dimg.DrawImage(img, dest, draw2d.NewTranslationMatrix(float64(i*w), float64(j*h)), draw.Over, draw2dimg.LinearFilter)

			icard++
		}
	}

	err = draw2dimg.SaveToPngFile(*args.outputFlag, dest)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func allCardsSameDimension(w, h int, path string, files []fs.DirEntry) error {
	for _, f := range files {
		fw, fh, err := getCardDimensions(path + f.Name())
		if err != nil {
			return err
		}

		if fw != w || fh != h {
			return fmt.Errorf("Dimensions do not match for file %s; Have: %d %d   Need: %d %d", f.Name(), fw, fh, w, h)
		}
	}
	return nil
}

func imageFromFilename(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}

func getCardDimensions(filename string) (int, int, error) {
	img, err := imageFromFilename(filename)

	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	return img.Bounds().Max.X - img.Bounds().Min.X, img.Bounds().Max.Y - img.Bounds().Min.Y, nil
}

func verify(args *Args) error {
	if *args.numColumnsOfCards < 1 {
		log.Fatal("Number of card columns must be >0")
	}
	if *args.numRowsOfCards < 1 {
		log.Fatal("Number of card rows must be >0")
	}

	return nil
}
