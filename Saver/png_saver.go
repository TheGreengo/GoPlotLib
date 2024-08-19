package png_saver

import (
    "os"
    "github.com/TheGreengo/GoPlotLib/Canvas"
)

// Constants
var HEAD[]byte         = []byte{0x89,0x50,0x4E,0x47,0x0D,0x0A,0x1A,0x0A}
var IEND_SIZE[]byte    = []byte{0x00,0x00,0x00,0x00}
var IEND_NAME[]byte    = []byte{0x49,0x45,0x4E,0x44}
var IEND_CRC[]byte     = []byte{0xAE,0x42,0x60,0x82}
const bit0 uint32      = 0xFF000000
const bit1 uint32      = 0x00FF0000
const bit2 uint32      = 0x0000FF00
const bit3 uint32      = 0x000000FF
// x^32+x^26+x^23+x^22+x^16+x^12+x^11+x^10+x^8+x^7+x^5+x^4+x^2+x+1
const CRC uint32     = 0x04C11DB7

func wrt_num(num uint32, file *os.File){
    bits := []byte{byte(num & bit0),byte(num & bit1), byte(num & bit2), byte(num & bit3)}
    file.Write(bits)
}

func MakeHeader(file *os.File) {
    file.Write(HEAD)
}

// The IHDR chunk must appear FIRST. It contains:
func MakeIHDR(file *os.File, c *canvas.Canvas) {
    // Width: from canvas
    wrt_num((*c).Width, file)
    // Height: from canvas
    wrt_num((*c).Height, file)
    // Bit depth: 8 bits for each color/alpha value
    file.Write([]byte{0x08})
    // Color type: r, g, b, alpha
    file.Write([]byte{0x06})
    // Compression method: only 0 works 
    file.Write([]byte{0x0})
    // Filter method: only 0 works 
    file.Write([]byte{0x0})
    // Interlace method: not interlaced
    file.Write([]byte{0x0})
}

func Compress() {
}

func CalcCRC() {
}

// each scanline (read: row) must have a filter type bit prepended
// 0 none, 1 sub, 2 up, 3 avg, 4 paeth
func MakeIDAT(f *os.File, l uint32, t uint32, c *canvas.Canvas, s int, e int) {
}

func CalcIDATs(file *os.File) int {
    return -1
}

func MakeIEND(file *os.File) {
    file.Write(IEND_SIZE)
    file.Write(IEND_NAME)
    file.Write(IEND_CRC)
}

func SaveCanvas(c *canvas.Canvas, s string) {
    // open file 
    file, err := os.Create(s) 
	if (err != nil) { panic(err) }
    
    MakeHeader(file)
    MakeIHDR(file, c)
    
    num := CalcIDATs(file)

    for i := 0; i < num; i++ {
        MakeIDAT(file, 4, 5, c, 6, 7)
    }
    
    MakeIEND(file)
}
