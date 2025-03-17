package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const BACK = "back.jpg"
const CARDS_IMAGE = "cards.jpg"
const DEFAULT_NUM_ROWS_CARDS = 10
const DEFAULT_NUM_COLUMNS_CARDS = 10
const MAX_IMAGE_WIDTH = 4096
const MAX_IMAGE_HEIGHT = 4096

const DEFAULT_CARD_WIDTH = 600
const DEFAULT_CARD_HEIGHT = 800
const DEFAULT_IMAGE_SOURCE = "."
const DEFAULT_TEST_IMAGE_DEST = "test_images"

var VERBOSE = false

// -b myback.png
// -d directory_of_images
// -o my_output_name.jpg

// Tabletop Simulator works best with "sheets" of up to 70 cards with 10 cards per row and 7 cards per column at 4090 x 4011 in maximum size for the overall sheet.
// https://www.stackup.org/post/creating-card-games-on-tabletop-simulator
// 10 width, 7 height
//
//
// https://entrogames.com/tabletop-simulator-ultimate-guide-absolute-beginners/
//

type Args struct {
	numColumnsOfCards      *int
	numRowsOfCards         *int
	testImagesWidth        *int
	testImagesHeight       *int
	backFlag               *string
	outputFlag             *string
	imageDirectoryFlag     *string
	testImageDirectoryFlag *string
	createTestImagesFlag   *bool
}

func main() {
	var args Args

	args.numColumnsOfCards = flag.Int("x", DEFAULT_NUM_COLUMNS_CARDS, "Number of card columns")
	args.numRowsOfCards = flag.Int("y", DEFAULT_NUM_ROWS_CARDS, "Number of card rows")

	args.backFlag = flag.String("b", BACK, "Name of card back jpeg file")
	args.outputFlag = flag.String("o", CARDS_IMAGE, "Name of jpeg output file (destination)")
	args.imageDirectoryFlag = flag.String("d", DEFAULT_IMAGE_SOURCE, "Name of image files directory (source)")
	args.testImageDirectoryFlag = flag.String("t", DEFAULT_TEST_IMAGE_DEST, "Name of test image files directory (destination). Creates if not exists. If exists, deletes all files in directory.")
	args.createTestImagesFlag = flag.Bool("C", false, "Create a set of test files")
	flag.BoolVar(&VERBOSE, "v", false, "Verbose output")
	args.testImagesWidth = flag.Int("w", DEFAULT_CARD_WIDTH, "Test image width")
	args.testImagesHeight = flag.Int("h", DEFAULT_CARD_HEIGHT, "Test image height")

	flag.Parse()

	fmt.Println(*args.backFlag)
	fmt.Println(*args.outputFlag)
	fmt.Println(*args.imageDirectoryFlag)
	fmt.Println(*args.createTestImagesFlag)

	// Create test files
	if *args.createTestImagesFlag {
		if err := createTestImages(*args.backFlag, *args.testImageDirectoryFlag, *args.testImagesWidth, *args.testImagesHeight); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		d2()
	} else {
		// Make single image page of directory of images
		if err := makeCardsPage(&args); err != nil {
			panic(err)
		}
	}

}

func verbose(mes string) {
	if VERBOSE {
		log.Println(mes)
	}
}
