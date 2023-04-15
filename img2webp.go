package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [input_file] [output_file]\n", os.Args[0])
	}

	inputFile :=os.Args[1]
	outputFile := os.Args[2]

	// input, err := os.Open(inputFile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// img, _, err := image.Decode(input)
	// if err != nil {
	// 	fmt.Println("Error decoding input file:",err)
	// 	os.Exit(1)
	// }

	img, err := imaging.Open(inputFile)
	if err != nil {
    log.Fatalf("Error decoding input file: %s", err)
	}


	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %s", err)
	}

	options := &webp.Options{
		Quality: 75,
	}
	err = webp.Encode(output, img, options)
	if err != nil {
		log.Fatalf("Error encoding image: %s", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}