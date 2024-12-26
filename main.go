package main

import (
    "fmt"
    "github.com/TheGreengo/GoPlotLib/Canvas"
    "github.com/TheGreengo/GoPlotLib/Saver"
    "os"
)

func main() {
    fmt.Println("GoPlotLib test started!")
    can := canvas.NewCanvas(4,4)
    fmt.Println(can)

    file, err := os.Create("temp.png")
    if (err != nil) { panic(err) }

    png_saver.MakeHeader(file)
    png_saver.MakeIEND(file)
    
    file.Close()
}
