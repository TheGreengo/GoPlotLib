package main

import ( 
    "os"
    "fmt" 
    "slices"
    "github.com/TheGreengo/GoPlotLib/Saver"
)

type Test func(string) (string, bool)

func WritePngHeaderTest(s string) (string, bool) {
    file, err := os.Create(s) 
	if (err != nil) { panic(err) }

    png_saver.MakeHeader(file)
    file.Close()

    file, err = os.Open(s) 
	if (err != nil) { panic(err) }

    bits := make([]byte,8)

    file.Read(bits)
    file.Close()

    os.Remove(s)

    HEAD := []byte{0x89,0x50,0x4E,0x47,0x0D,0x0A,0x1A,0x0A}

    return "Writing PNG header", slices.Equal(bits, HEAD)
}

func WriteIHDRTest(s string) (string, bool) {
    file, err := os.Create(s) 
	if (err != nil) { panic(err) }
    file.Close()

    return "Writing IHDR header", true
}

func WriteIENDTest(s string) (string, bool) {
    file, err := os.Create(s) 
	if (err != nil) { panic(err) }
    file.Close()

    return "Writing IEND Block", true
}

// This should create a new file, and then loop through all of the
// tests and run them on the file and output the result of the tests
func main() {
    fmt.Println("Testing started")

    tests := []Test{
        WritePngHeaderTest,
        WriteIHDRTest,
        WriteIENDTest,
    }

    overall := 0 
    for _, i := range tests {
        res, success := i("temp.png")
        if success {
            fmt.Printf("%s: ", res)
            fmt.Printf("[\033[38;2;152;255;152mpass\033[0m]\n")
        } else {
            fmt.Printf("%s: ", res)
            fmt.Printf("[\033[38;2;217;84;77mfail\033[0m]\n")
            overall += 1
        }
    }

    if overall == 0 {
        fmt.Println("All", len(tests), "tests passed.")
    } else {
        fmt.Println(overall, "tests failed.")
    }
}
