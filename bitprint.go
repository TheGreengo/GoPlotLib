package main

import ( "fmt";"os";"flag" )

func PrintBits(fil string) {
    file, err := os.Open(fil) 
	if (err != nil) { 
        fmt.Println("Issues opening", fil)
        panic(err) 
    }
    
    bits := make([]byte,8)
    out  := make([][]string,2)

    out[0] = make([]string,8)
    out[1] = make([]string,8)

    for {
		n, err := file.Read(bits) 

		if (err != nil) {
			break
		}

        fmt.Println("---------------------------------------------------------------------------------------")
        for i := 0; i < n; i++ {
            fmt.Printf("%10d ", bits[i])
        }
        fmt.Printf("\n")

        for i := 0; i < n; i++ {
            if (bits[i] != 10) {
                fmt.Printf("%10c ", bits[i])
            } else {
                fmt.Printf("%10s ", "\\n")
            }
        }
        fmt.Printf("\n")

        for i := 0; i < n; i++ {
            fmt.Printf("| ")
            for j := 0; j < 8; j ++ {
                fmt.Printf("%d", ((128 >> j) & bits[i]) >> (7-j))
            }
            fmt.Printf(" ")
        }
        fmt.Printf("\n")
    }
}

func main() {
	file_ptr := flag.String("f","", "The name of the file to be encrypted")

	flag.Parse()

    PrintBits(*file_ptr)
}
