package MaekSmol

/*
* So..... the steps are thus:
* - get a count the frequency of all of the bytes
* - make a tree with all of the things
* - use this to make the encodings
* - write a copy of the file with the encoded values
* - figure out a way to un-encode the stuff
*/

import (
    "os"
)

type ByteFreq struct {
    val byte
    freq uint
}
 
// This should do the whole thring
func HoffmanEncode() {
}

// This should go through 
func GetTree() {
}

// This should go through the file, count every symbol, and return 
func CountSymbols(name string) []ByteFreq {
    var bytes make([]ByteFreq, 256)

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
        
        for i = 0; i < n; i++ {
            bytes[int(bits[i])]
        }
    }

    file.Close()

    return counts
}
