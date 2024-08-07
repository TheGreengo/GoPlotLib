package MaekSmol

/*
* So..... the steps are thus:
* - get a count the frequency of all of the bytes
* - make a tree with all of the things
* - use this to make the encodings
* - write a copy of the file with the encoded values
* - figure out a way to un-encode the stuff
*/

import ( "os";"fmt";"sort";"container/heap" )

type ByteFreq struct {
    val  byte
    freq uint
}
 
type ByteNode struct {
    freq  uint
    leaf  bool
    val   byte
    left  *ByteNode
    right *ByteNode
}

func printTree(bn *ByteNode, curr string) {
    if (*bn).leaf {
        fmt.Printf("%4c %10s %10d\n", (*bn).val,curr,(*bn).freq)
    }
    if (*bn).left != nil {
        printTree((*bn).left, curr + "0")
    }
    if (*bn).right != nil {
        printTree((*bn).right, curr + "1")
    }
}
// This should do the whole thring
// Note: we could do this a faster way
func HuffmanEncode(name string) {
    counts := CountSymbols(name)

    // Here we need to 
    ordered := make(PrioQue, len(counts)) 
    
    for i,bf := range counts {
        ordered[i] = ByteNode{
            freq: bf.freq,
            leaf: true,
            val: bf.val,
            left: nil,
            right: nil,
        }
    }

    heap.Init(&ordered)
    heap.Push(&ordered, ByteNode{freq: uint(4), leaf: false, val: byte(0), left: nil, right: nil})

    fmt.Println("Ordered", ordered)
    for len(ordered) > 1 {
        l1 := heap.Pop(&ordered).(ByteNode)
        l2 := heap.Pop(&ordered).(ByteNode)

        nn := ByteNode{
            freq: l1.freq + l2.freq,
            leaf: false,
            val: byte(0),
            left: &l2,
            right: &l1,
        }

        heap.Push(&ordered, nn)
    }
    final := heap.Pop(&ordered).(ByteNode)
    printTree(&final,"")
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
