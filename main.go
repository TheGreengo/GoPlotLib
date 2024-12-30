package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
)

/*
Fun ideas:
- draw axis
- draw lines
- set color
*/

// t = thickness of the line
func draw_line(img *image.RGBA, x1 int, y1 int, x2 int, y2 int, t int, col color.RGBA) {
    // so somehow we need to get all of the
    steps := 500
    dx := float64(x2-x1) / float64(steps)
    dy := float64(y2-y1) / float64(steps)

    for j := 0; j < t; j++ {
        x := float64(x1)
        y := float64(y1)
        for i := 0; i < steps; i++ {
            x += dx
            y += dy
            img.Set(int(x)+j, int(y), color.RGBA{0, 255, 0, 255})
        }
    }
}

func main() {
    farve := color.RGBA{255, 255, 255, 255}
	// Create a new image with a specific width and height
	width  := 640
	height := 480
	img    := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with red color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, farve)
		}
	}

    draw_line(img, 10, 10, 100, 400, 4, farve)

	// Create a file to save the image
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image as PNG and write to file
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
