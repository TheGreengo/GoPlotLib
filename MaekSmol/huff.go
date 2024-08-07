package MaekSmol

/*
* So..... the steps are thus:
* - get a count the frequency of all of the bytes
* - make a tree with all of the things
* - use this to make the encodings
* - write a copy of the file with the encoded values
* - figure out a way to un-encode the stuff
*/

import ( "os";"fmt";"sort" )

type ByteFreq struct {
    val byte
    freq uint
}
 
// This should do the whole thring
func HuffmanEncode(name string) {
    counts := CountSymbols(name)

    // Here we need to 
    ordered := make(PrioQue, len(counts)) 
    
    for i,thing := range counts {
        ordered[i] = &ByteFreq{
            val: thing.val,
            freq: thing.freq,
        }
    }

    fmt.Println("Ordered",ordered)
    for len(ordered) > 0 {
        thing := ordered.Pop()
        fmt.Println(thing)
    }
}

// This should go through 
func GetTree() {
}

// This should go through the file, count every symbol, and return 
func CountSymbols(name string) []ByteFreq {
    bytes := make([]ByteFreq, 256)

    var i uint
    for i = 0; i < 256; i++ {
        bytes[i] = ByteFreq{val: byte(i), freq: 0}
    }

    file, err := os.Open(name)
	if (err != nil) { panic(err) }
	
    bits := make([]byte,1000000)
    for {
		n, err := file.Read(bits) 
		
		if (err != nil) { break }
        
        for i = 0; i < uint(n); i++ {
            bytes[int(bits[i])].freq += 1
        }
    }
    
    file.Close()
    
    sort.Slice(bytes, func(i, j int) bool {
        return bytes[i].freq > bytes[j].freq
    })

    low := uint(1)
    for i = 0; i < 256; i++ {
        low = bytes[i].freq
        if low == 0 { break }
    }

    bytes = bytes[0:i]

    return bytes 
}
