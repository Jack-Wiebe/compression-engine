package main

import (
	"bufio"
	"compression-engine/util"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//test.txt -> 2536 -> 1204
var code_map map[string]string

func main() {

	//data structures
	var input string
	var output string
	var bit_array []rune
	freq_table := map[byte]int {}
	code_map = make(map[string]string)
	queue := make(util.PriorityQueue, 0)

	//initialize flag arguments
	filePath := flag.String("file", "test/default.txt", "input file")
	flag.Parse()
	fmt.Println(*filePath)

	//open and read file
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	//
	//HUFFMAN ENCODER
	//

	//scan data and create Freq Table
	for scanner.Scan() {
		byteData := scanner.Bytes()[0]
		input+=string(byteData);

		bits := make([]rune, 8)
		for i := 0; i < 8; i++ {
			// Shift the byte right by i positions and mask with 1 to get the i-th bit
			nextBit := rune((byteData >> i) & 1)
			bits[7-i] = nextBit
			bit_array = append(bit_array, nextBit)
		}

		if _, ok := freq_table[byteData]; ok {
			freq_table[byteData]++
			fmt.Printf("char %c freq: %d\n",byteData, freq_table[byteData])
		} else {
			fmt.Printf("adding %c to map\n", byteData)
			freq_table[byteData] = 1
		}
		// Print the bits
		fmt.Printf("Char: %c\n", byteData)
		fmt.Printf("Bits: %v\n", bits)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	fmt.Println()
	fmt.Println(bit_array)
	fmt.Println()
	fmt.Println("Frequency Table:")
	util.PrintMap(freq_table)
	fmt.Println()

	//populate priority queue
	heap.Init(&queue)
	for char,freq := range freq_table{
		heap.Push(&queue, &util.Node{Value: string(char), Priority: freq})
	}
	fmt.Println("Priority Queue:")
	queue.Print()

	//populate binary tree
	for queue.Len() > 1 {
		item1 := heap.Pop(&queue).(*util.Node)
		item2 := heap.Pop(&queue).(*util.Node)
		node := &util.Node{Value: item1.Value+item2.Value, Priority: item1.Priority+item2.Priority}
		node.Right = item1
		node.Left = item2
		queue.Push(node)
	}
	root := heap.Pop(&queue).(*util.Node)
	fmt.Println("Binary Tree:")
	root.PreOrderTraversal()
	fmt.Println()

	//generate code map and output
	fmt.Println("Code Map")
	GetCodes(root, "")
	fmt.Println()
	fmt.Println(input)
	for _,char := range input{
		output+=code_map[string(char)]
	}
	fmt.Println(output)

	//last 3 bits encode significant bits in last 2 bytes
	remainder := len(output)%8
	padding := ""
	if remainder != 0 {
		fmt.Printf("additional bits: %d\n",remainder)
		for i := 0; i < 13 - remainder; i++ {
			padding+="0"
		}
		fmt.Printf("padding: %s\n", padding)
		binaryCount := strconv.FormatInt(int64(remainder), 2)
		significant_bits := fmt.Sprintf("%03s", binaryCount)
		fmt.Printf("suffix: %s\n", significant_bits)
		output+=padding+significant_bits;
	}

	//convert binary string to byte array
	byteArray, err := util.ConvertToByteArray(output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println()
	fmt.Println(byteArray)

	//output to binary file
	outputFile, err := os.Create("test/output.bin")
	if err != nil {
		fmt.Printf("failed to create file: %v", err)
	}
	defer outputFile.Close()
	_, err = outputFile.Write(byteArray)
	if err != nil {
		fmt.Printf("failed to write to file: %v", err)
	}

	//
	//HUFFMAN DECODER
	//

	// Read the entire file into a byte slice
	data, err := os.ReadFile("test/output.bin")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert the byte slice to a string
	content := util.ConvertToBinaryString(data)
	fmt.Println(content)

	//read suffix for significant bits
	suffix := content[len(content)-3:]
	significant_bits,_ := strconv.ParseInt(suffix, 2, 64)

	//decode binary content
	decoded_string := ""
	current_symbol := ""
	for i,char := range content{
		if i >= len(content) - 16 + int(significant_bits){
			fmt.Println("EOF reached... ---> ", current_symbol)
			break;
		}
		current_symbol += string(char)
		value, exists := util.FindKeyByValue(code_map, current_symbol)
		if(exists){
			decoded_string += value
			current_symbol = ""
		}
	}
	fmt.Println(decoded_string)

	//Context-Tree Weighting prediction loop

	//compression engine:
	//parser
	//context tree bank
	//output to file
	//ingestion engine
}



func GetCodes(n *util.Node, code string) {
	if n == nil {
		return
	}

	GetCodes(n.Left, code+"1")
	GetCodes(n.Right, code+"0")

	if(n.Left == nil && n.Right == nil){
		fmt.Printf("%s: %s\n", n.Value, code)
		code_map[n.Value]=code
	}
}

