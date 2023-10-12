package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/disintegration/gift"
	"github.com/makeworld-the-better-one/dither/v2"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

var palette = color.Palette{
	color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255}, // Black, White,
	color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{255, 0, 0, 255}, // Green, Blue, Red
	color.RGBA{255, 255, 0, 255}, color.RGBA{255, 128, 0, 255}, // Yellow, Orange
}

var inFile string
var outFile string
var brightness float64
var contrast float64
var strength float64

var input image.Image
var tweaked *image.RGBA
var output image.Image

func init() {
	flag.StringVar(&inFile, "i", "", "Input filename")
	flag.Float64Var(&brightness, "b", 0, "Adjust brightness -100 to 100. 0 for none")
	flag.Float64Var(&contrast, "c", 0, "Adjust contrast -100 to 100. 0 for none")
	flag.Float64Var(&strength, "d", 1, "Dithering strength 0.0 to 1.0")
	flag.Parse()

	if _, err := os.Stat(inFile); err != nil {
		_, bin := filepath.Split(os.Args[0])
		fmt.Printf("Input file not found.\n\n")
		fmt.Printf("Usage: %s -i input.jpg\n", bin)
		flag.PrintDefaults()
		os.Exit(1)
	}

	outFiles := strings.Split(inFile, ".")
	outFiles = outFiles[:len(outFiles)-1]
	outFile = strings.Join(outFiles, ".")
	outFile += "-out.bmp"
}

func main() {
	fmt.Printf("Config:\n")
	fmt.Printf("- input file: %s\n", inFile)
	fmt.Printf("- output file: %s\n", outFile)
	fmt.Printf("- brightness: %f\n", brightness)
	fmt.Printf("- contrast: %f\n", contrast)
	fmt.Printf("- dither strength: %f\n", strength)

	// read input
	f, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	input, _, err = image.Decode(bytes.NewReader(f))
	if err != nil {
		fmt.Printf("Error decoding input file: %v", err)
		os.Exit(1)
	}

	// adjust brightness/contrast
	g := gift.New(gift.Brightness(float32(brightness)), gift.Contrast(float32(contrast)))
	tweaked = image.NewRGBA(g.Bounds(input.Bounds()))
	g.Draw(tweaked, input)

	// dither image
	d := dither.NewDitherer(palette)
	d.Matrix = dither.ErrorDiffusionStrength(dither.FloydSteinberg, float32(strength))
	output = d.Dither(tweaked)

	// write output
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v", err)
		os.Exit(1)
	}

	err = bmp.Encode(out, output)
	if err != nil {
		fmt.Printf("Error encoding output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Done\n")
}
