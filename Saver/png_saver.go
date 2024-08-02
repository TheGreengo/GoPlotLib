package png_saver

import (
    "os"
    "github.com/TheGreengo/GoPlotLib/Canvas"
)

// Constants
var HEAD      = [...]byte{0x89,0x50,0x4E,0x47,0x0D,0x0A,0x1A,0x0A}
var IEND_SIZE = [...]byte{0x00,0x00,0x00,0x00}
var IEND_NAME = [...]byte{0x49,0x45,0x4E,0x44}
var IEND_CRC  = [...]byte{0xAE,0x42,0x60,0x82}

func MakeHeader(file * File) {
    file.write(HEAD)
}

// The IHDR chunk must appear FIRST. It contains:
//   Width:              4 bytes
//   Height:             4 bytes
//   Bit depth:          1 byte
//   Color type:         1 byte
//   Compression method: 1 byte
//   Filter method:      1 byte
//   Interlace method:   1 byte
func MakeIHDR(file * File, c *canvas) {
}

func CRC() {
}

func MakeIDAT(l uint_32, t uint_32, c *Canvas, s int, e int) {
}

func CalcIDATs(file * File) int {
    return -1
}

func MakeIEND(file * File) {
    file.write(IEND_SIZE)
    file.write(IEND_NAME)
    file.write(IEND_CRC)
}

func SaveCanvas(c *Canvas, s string) {
    // open file 
    file, err := os.Create(s) 
	if (err != nil) {
		panic(errr)
	}
    
    MakeHeader(file)
    MakeIHDR(file)
    
    num := CalcIDATs()

    for i := 0; i < num; i++ {
        MakeIDAT()
    }
    
    MakeIEND(file)
}
