package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func main() {
	var quality int
	flag.IntVar(&quality, "q", 75, "Quality value for WebP compression (0-100)")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Printf("Usage: %s [input_file] [output_file] [flags]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile := flag.Args()[0]
	outputFile := flag.Args()[1]
	// if flag.Arg(3) != "" {
	// 	quality, _ = strconv.Atoi(flag.Arg(3))
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
		Quality: float32(quality),
	}
	err = webp.Encode(output, img, options)
	if err != nil {
		log.Fatalf("Error encoding image: %s", err)
	}

	fmt.Printf("Successfully converted %s to %s with quality %d\n", inputFile, outputFile, quality)
}
