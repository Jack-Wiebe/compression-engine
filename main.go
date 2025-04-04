package main

import (
	"compression-engine/huffman"
	"compression-engine/pass_extraction"
	"compression-engine/runlength"
	"flag"
	"fmt"
	"log"
	"os"
)

//test.txt -> 2536 -> 1204

func main(){

	Runlength_Huffman_Test()

	pass_extraction.Test()

}

func Runlength_Huffman_Test() {

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

}

