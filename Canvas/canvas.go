package canvas

const RED int = 0
const GREEN int = 1
const BLUE int = 2
const ALPHA int = 3

type ColorByte struct {
    data [4]byte
}

func (cb *ColorByte) set_date(r byte, g byte, b byte, a byte) {
    cb.data[RED] = r
    cb.data[GREEN] = g
    cb.data[BLUE] = b
    cb.data[ALPHA] = a
}

func (cb *ColorByte) get_red() byte {
    return cb.data[RED]
}

func (cb *ColorByte) get_green() byte {
    return cb.data[GREEN]
}

func (cb *ColorByte) get_blue() byte {
    return cb.data[BLUE]
}

func (cb *ColorByte) get_alpha() byte {
    return cb.data[ALPHA]
}

type Canvas struct {
    height int
    width int
    colors [][]ColorByte
}

func NewCanvas(w int, h int) Canvas {
    dat := make([][]ColorByte, h, h)
    for i := 0; i < h; i++ {
        dat[i] = make([]ColorByte, w, w)
        for j := 0; j < w; j++ {
            d := [4]byte{0,0,0,0}
            dat[i][j] = ColorByte{ data: d}
        }
    }
    return Canvas{height: h, width: w, colors: dat}
}
