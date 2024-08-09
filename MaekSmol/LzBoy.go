package MaekSmol

import ( "os" )

// This should do the LZ77 algorithm analysis stuff and

// some quick notes:
// - the length is 3-258 bytes (represented as one byte)
// - the offset is 1-32,768

// I don't understand this, but there's going to like be 285 symbols or ngt 

const LEN_OFFSET  uint = 3
const WINDEHR_LEN uint = 32678

func LZ77() {
    // make window
    // pull in first character
    // loop through, checking if the window contains a substring of more
    // than 2 characters
}

// This should be the helper function to find if something is in 
// the sliding window
func CheckWindehr() {
}
