package main

import (
	"fmt"
	"image"
	"os"

	"github.com/chai2010/webp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [input_file] [output_file]\n", os.Args[0])
	}

	inputFile :=os.Args[1]
	outputFile := os.Args[2]

	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img, _, err := image.Decode(input)
	if err != nil {
		fmt.Println("Error decoding input file:",err)
		os.Exit(1)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:",err)
		os.Exit(1)
	}

	options := &webp.Options{
		Quality: 80,
	}
	err = webp.Encode(output, img, options)
	if err != nil {
		fmt.Println("Error encoding image:",err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}