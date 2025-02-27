package main

import "flag"
import "fmt"
import "os"

const BACK = "back.jpg"
const OUTPUT = "cards.jpg"
const DEFAULT_CARD_WIDTH = 64
const DEFAULT_CARD_HEIGHT = 160

// -b myback.png
// -d directory_of_images
// -o my_output_name.jpg

func main() {
	var backFlag = flag.String("b", BACK, "Name of card back jpeg file")
	var outputFlag = flag.String("o", OUTPUT, "Name of jpeg output file (destination)")
	var imageDirectoryFlag = flag.String("d", ".", "Name of image files directory (source)")
	var createTestImagesFlag = flag.Bool("C", false, "Create a set of test files")

	flag.Parse()

	fmt.Println(*backFlag)
	fmt.Println(*outputFlag)
	fmt.Println(*imageDirectoryFlag)

	if *createTestImagesFlag {
		if err := createTestImages(*backFlag, *imageDirectoryFlag); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		d2()
	}

}
