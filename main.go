package main

import (
    "fmt"
    "github.com/TheGreengo/GoPlotLib/Canvas"
)

func main() {
    fmt.Println("GoPlotLib test started!")
    can := canvas.NewCanvas(4,4)
    fmt.Println(can)
}
