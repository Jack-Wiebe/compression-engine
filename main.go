package main

import (
	"compression-engine/huffman"
	"compression-engine/lz"
	"compression-engine/runlength"
	"flag"
	"fmt"
	"log"
	"os"
)

//test.txt -> 2536 -> 1204

func main(){

	//pass_extraction.Test()

	Encoding_Test()
}

func Encoding_Test() {

	filePath := flag.String("file", "test/default.txt", "input file")
	flag.Parse()
	fmt.Println(*filePath)

	//open and read file
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	huffman.Encode_Huffman(file)
	huffman.Decode_Huffman()

	runlength.Encode_Runlength(file)
	runlength.Decode_Runlength()


  lz.Encode_LZSS(file)
	//lz.Encode_LZ77(file)

}

