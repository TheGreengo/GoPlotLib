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
- label graph
- label axes
- make points floats
*/

var (
    RED     color.RGBA = color.RGBA{ 255,   0,   0, 255 } // done
    ORANGE  color.RGBA = color.RGBA{ 255, 165,   0, 255 } // done
    YELLOW  color.RGBA = color.RGBA{ 255, 255,   0, 255 } // done
    FORREST color.RGBA = color.RGBA{  34, 139,  34, 255 } // done
    GREEN   color.RGBA = color.RGBA{ 175, 255,  47, 255 } // done
    BLUE    color.RGBA = color.RGBA{   0,   0, 255, 255 } // done
    PURPLE  color.RGBA = color.RGBA{ 148,   0, 211, 255 } // done
    MAGENTA color.RGBA = color.RGBA{ 255,   0, 255, 255 } // done
    WHITE   color.RGBA = color.RGBA{ 255, 255, 255, 255 } // done
    GRAY    color.RGBA = color.RGBA{ 211, 211, 211, 255 } // done
    BLACK   color.RGBA = color.RGBA{   0,   0,   0, 255 } // done
    BROWN   color.RGBA = color.RGBA{ 139,  69,  19, 255 } // done
    PINK    color.RGBA = color.RGBA{ 255,  20, 147, 255 } // done
    S_BLUE  color.RGBA = color.RGBA{   0, 191, 255, 255 } // done
    SEETHR  color.RGBA = color.RGBA{ 255, 255, 255, 255 } // done
)

func draw_rect() {
}

func draw_circle() {
}

func draw_rrect() {
}

func draw_slice() {
}

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
            img.Set(int(x)+j, int(y), col)
        }
    }
}

func make_plot(img *image.RGBA) {
    //dx := img.Bounds().Dx()
    dy := img.Bounds().Dy()
    draw_line(img, 5, 5, 5, dy - 5, 2, color.RGBA{0, 0, 0, 255})
}

func main() {
	// Create a new image with a specific width and height
	width  := 640
	height := 480
	img    := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with red color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, WHITE)
		}
	}

    draw_line(img, 100, 100, 600, 200, 4, GREEN)

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
