package main

import (
	"./resize"
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
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := toRect(100, 100)

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
	fl, err := os.OpenFile("out.jpg", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("couldn't write", err)
		return
	}
	defer fl.Close()
	jpeg.Encode(fl, outputImage, nil)
}
