package Canvas

const RED int = 0
const GREEN int = 1
const BLUE int = 2
const ALPHA int = 3

type Canvas struct {
    height uint32
    width uint32
    colors [][][4]byte
}

func NewCanvas(w uint32, h uint32) Canvas {
    var i, j uint32

    dat := make([][][4]byte, h)

    for i = 0; i < h; i++ {
        dat[i] = make([][4]byte, w)

        for j = 0; j < w; j++ {
            d := [4]byte{0x0,0x0,0x0,0x0}
            dat[i][j] = d
        }
    }
    return Canvas{height: h, width: w, colors: dat}
}

func (c *Canvas) set_pix(row int, col int, r int, g int, b int, a int) {
    c.colors[row][col][RED]   = byte(r)
    c.colors[row][col][GREEN] = byte(g)
    c.colors[row][col][BLUE]  = byte(b)
    c.colors[row][col][ALPHA] = byte(a)
}
