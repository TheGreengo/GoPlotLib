package MaekSmol

import (
    "os"
)

// This should take a file, and then 
func HoffmanEncode() {
}

// This should go through 
func getTree() {
}

// This should go through the file, count every symbol, and return 
func CountSymbols(name string) map[byte]int {
    counts := make(map[byte]int)

    file, err := os.Open(name)

	if (err != nil) { panic(err) }
	
    // Then we loop through all the bytes, and make a map with the counts for
    // all the bytes
    bits := make([]byte,1000000)
    for {
		n, err := file.Read(bits) 
		
		if (err != nil) { break }
        
        for i := 0; i < n; i++ {
            _, cont := counts[bits[i]]
            if (!cont) {
                counts[bits[i]] = 1 
            } else {
                counts[bits[i]] += 1
            }

        }
    }
    return counts
}
