package main

import (
	"./resize"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func toRect(w, h int) image.Rectangle {
	return image.Rect(0, 0, w, h)
}

func isSquare(r image.Rectangle) bool {
	return r.Dx() == r.Dy()
}
func isTall(r image.Rectangle) bool {
	return r.Dy() > r.Dx()
}

func isWide(r image.Rectangle) bool {
	return r.Dx() > r.Dy()
}

func main() {
	var source string
	flag.StringVar(&source,"source", "./test.jpg", "image to read")
	var dest string
	flag.StringVar(&dest, "dest", "./out.jpg", "image to output")
	
	var width, height int

	flag.IntVar(&width, "width", 100, "width to resize to")
	flag.IntVar(&height, "height", 100, "height to resize to")

  flag.Parse()
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := toRect(width, height)

	// Decode the image.
	m, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("error decoding image\n")
		log.Fatal(err)
	}
	bounds := m.Bounds()
	if isWide(bounds) {
		fmt.Println("wider than taller")
	}
	if isTall(bounds) {
		fmt.Println("taller than wider %d %d", bounds.Dx(), bounds.Dy())
	}
	if isSquare(bounds) {
		fmt.Println("square")
	}
	if isSquare(s) {
		fmt.Println("dimensions are square")
	}
	outputImage := resize.Resize(m, bounds, s.Dx(), s.Dy())
	outBounds := outputImage.Bounds()
	fmt.Printf("%q\n", outBounds)
	fl, err := os.OpenFile(dest, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("couldn't write", err)
		return
	}
	defer fl.Close()
	jpeg.Encode(fl, outputImage, nil)
}
