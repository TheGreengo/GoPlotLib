# GoPlotLib
My own stab at making a library for generating plots

## Deflate Algorithm

### General Flow

First, the LZ77 encoding algorithm is run on the binary file. The end result of that should be a file with a set of bytes with the raw values as well as the lengths and pointers encoded as well. 

Then the Huffman encoding takes place, encoding for 0-288 with the following breakdown

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
