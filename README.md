# GoPlotLib
My own stab at making a library for generating plots

## Deflate Algorithm

### General Flow

First, the LZ77 encoding algorithm is run on the binary file. The end result of that should be a file with a set of bytes with the raw values as well as the lengths and pointers encoded as well. 

Then the Huffman encoding takes place, encoding for 0-288 with the following breakdown

- data other than the trees is packed from least to most significant bit
- the Huffman trees are packed most to least significant bit

|Number|Value|
|----|----|
|0-255|These values are used to encode raw bytes (i.e. literals not representing a reference back to an earlier string|
|256|The end of a block. This means either that the next block is starting, or, if this block started with the last block bit, the end of the file.|
|257-285|An 8-bit encoding of the length of the following match|
|287-288|Reserved for future purposes. Shouldn't be used.|

There will also be a separate tree for encoding the distances. This results in the final form of each block as being 

| ID | Symbol Tree | Distance Tree | Compressed Data|
|--|--|--|--|



### LZ77

### Huffman Encoding

Using A = 10, B = 0, C = 110, D = 111

Codes must follow two rules:
- All codes are lexographically ordered so that short codes are lexographically before long ones. I.e. if we have A represented by 2 bits and B represented by 1 bit, B should be 0 and A should be 10 because in a lexographic ordering, 0 is before 10. In this case, you can think about lexographic meaning that if you went bit by bit, you would find that if you sorted characters by which ones had the lowest (i.e. 0 instead of 1) first bit, then lowest second bit if the first bits match, and so on. In the example given, you would look at the first bit of B (which is 0) and the first bit of A (which is 1), you would put B first. When comparing A to C or D, you would see their first bits "match" so you would look at the second bit and see that A comes before C and D since its second bit is 0.
- All codes of equal length are ordered lexographically in the same ordering as the symbols they represent. In the case of our example, this means that since C and D are represented by 3 bits starting with `11`, we would determine the third bit by ordering them alphabetically, since they represent letters, meaning that C = `110` and D = `111`

## Todos
- get the DEFLATE algorith working
    - get Huffman encodings to work into binary
    - figure out the lz77 algorithm
    - get that to work with binary
    - figure out the format required by deflate
- figure out what format requirements .png imposes
- get the CRC algorithm working
    - get the CRC codes to be reproducable
    - get the CRC setter to work with setting the final field in a chunk
- get the bits chunks actually set properly
- make a .png
