package png_saver

import (
    "os"
    "github.com/TheGreengo/GoPlotLib/Canvas"
)

// Constants

const HEAD uint_64 = 0x89504E470D0A1A0A

const IEND_SIZE uint_32 = 0x00000000
const IEND_NAME uint_32 = 0x49454E44
const IEND_CRC uint_32 = 0xAE426082

func CRC() {
}

// ALL Chunks
//  - length
//  - type/name
//  - data
//  - CRC

func MakeHeader() {
}

// The IHDR chunk must appear FIRST. It contains:
//   Width:              4 bytes
//   Height:             4 bytes
//   Bit depth:          1 byte
//   Color type:         1 byte
//   Compression method: 1 byte
//   Filter method:      1 byte
//   Interlace method:   1 byte
func MakeIHDR() {
}

func MakeIDAT(l uint_32, t uint_32, d *Canvas, s int, e int) {
    var CRC uint_32 = 0x00000000
}

func CalcIDATs() int {
    return -1
}

func MakeIEND() {
}

func SaveCanvas(c *Canvas, f string) {
    // open file 
    file, errr := os.Create(*out_ptr) 
	if (errr != nil) {
		panic(errr)
	}
    
    MakeHeader()
    MakeIHDR()
    
    num := CalcIDATs()

    for i := 0; i < num; i++ {
        MakeIDAT()
    }
    
    MakeIEND()
}
