package main

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
)

func makeCardsPage(args *Args) error {
	log.Println("TODO: func makeCardsPage(args *Args) {")
	log.Println(*args.imageDirectoryFlag)
	if err := verify(args); err != nil {
		return err
	}

	if _, err := os.Stat(*args.imageDirectoryFlag); os.IsNotExist(err) {
		return errors.New("Image files directory (source) " + *args.imageDirectoryFlag + " does not exist.")

	}

	root := os.DirFS(*args.imageDirectoryFlag)
	files, err := fs.ReadDir(root, ".")

	if err != nil {
		log.Fatal(err)
	}

	w, h, err := getCardDimensions(*args.imageDirectoryFlag + "/" + files[0].Name())
	log.Println(w, h)

	if err != nil {
		return err
	}

	for _, f := range files {
		fw, fh, err := getCardDimensions(*args.imageDirectoryFlag + "/" + f.Name())
		log.Println(*args.imageDirectoryFlag+"/"+f.Name(), fw, fh, w, h)
		if err != nil {
			return err
		}
		log.Println(fw != w || fh != h)
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
