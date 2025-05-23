package lz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EncodingType int

const (
    LZ77 EncodingType = iota  // 0
    LZSS                			// 1
)

const window_length = 8

//TODO: FIX STRING ENCODING
func Encode_LZ77(file *os.File) {


	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking:", err)
		return
	}

	fmt.Println("LZ77 Encoding", file.Name())
	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n") // Preserve line breaks
	}
	s := builder.String()
	fmt.Println(s)


	input := []rune(s)
	output := input[0]
	characters_processed := window_length/2


	//create buffer and populate with initial data
	buffer := make([]rune, window_length)
	copy(buffer,append([]rune(strings.Repeat(string(input[0]), window_length/2)), input[0:window_length/2]...) )

	fmt.Println("buffer:",buffer)
	fmt.Println(string(buffer))
	fmt.Println("output:",string(output))
	encoded_string := string(output)

	//Begin loop
	for (len(input) > 0) {

		//find next best match in buffer
		result, offset := findMatch(buffer)

		//encode match to output and move buffer forward
		encoded_string += result
		fmt.Println()
		fmt.Println("input:", string(input))
		if window_length -(offset+1) <= len(input){ //check if offset exceeds length of remaining input
			//todo: remaining input is skipped
			//oder is flipped - buffer needs to be shifted by offset first
			// then grab next characters from input and add to buffer
			// then pop charaters from front of input
			input = input[characters_processed:]
			characters_processed = offset+1
			copy(buffer,append(buffer[offset+1:], input[0:window_length-(offset+1)]...)) //TODO: redo this

			// copy(buffer,append(buffer[offset+1:], input[0:offset]...))
			// input = input[offset+1:]
		}else{
			input =  make([]rune, 0)
		}

		fmt.Println("input:", string(input))
		fmt.Println("encoded string:", encoded_string)
		fmt.Println("buffer:", string(buffer))
		fmt.Println()
	}
}

func Encode_LZSS(file *os.File) {

	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking:", err)
		return
	}

	fmt.Println("LZSS Encoding", file.Name())
	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n") // Preserve line breaks
	}
	s := builder.String()
	fmt.Println(s)


	//SET UP CONSTANTS AND VARIABLES
	input := []rune(s)
	output := ""
	SMALLEST_MATCH := 7
	FRONT_BUFFER_SIZE := 1024
	BACK_BUFFER_SIZE := 1024

	if len(input) < FRONT_BUFFER_SIZE {
		FRONT_BUFFER_SIZE = len(input)
	}
	if len(input) < BACK_BUFFER_SIZE {
		BACK_BUFFER_SIZE = len(input)
	}

	//instantiate buffer
	buffer := make([]rune, BACK_BUFFER_SIZE+FRONT_BUFFER_SIZE)
	copy(buffer,append([]rune(strings.Repeat(string('\x00'), FRONT_BUFFER_SIZE)), input[0:BACK_BUFFER_SIZE]...))
	input = input[BACK_BUFFER_SIZE:]

	//instantiate windows
	front := buffer[:FRONT_BUFFER_SIZE]
	back  := buffer[BACK_BUFFER_SIZE:]


	fmt.Println("Initialize....")
	fmt.Println("input:",input, string(input))
	fmt.Println("buffer:",buffer, string(buffer))
	fmt.Println("front:",front, string(front))
	fmt.Println("back:",back, string(back))


	for len(back) > 0{

		length := 0
		offset := 0
		match := ""
		fmt.Println()
		fmt.Println("buffer:", string(front),string('|'), string(back))

		for i := range(front){
			if buffer[i] == back[0]{
				for j :=range(back){
					if buffer[i+j] == back[j]{
						if (j > length){
							offset = i
							length = j
							match = string(buffer[offset:offset+length])
						}
					}else{
						break
					}
				}
			}
		}

		shift := length+1

		//encode string
		if(length < SMALLEST_MATCH){
			output += string(back[0])
			shift = 1
		}else{
			fmt.Println("adding length and offset to output:",length, offset)
			output += "<"
			output += strconv.Itoa(FRONT_BUFFER_SIZE-offset)
			output += ":"
			output += strconv.Itoa(length+1)
			output += ">"
		}



		//move by offset
		fmt.Println("shift buffer by:", length+1)
		if shift > len(input){
			//reached end of output
			fmt.Println("pulling last of input", string(input))
			//buffer = append(buffer[shift:],input...)
			buffer = buffer[shift:]
		}else{
			fmt.Println("shifting input into buffer", string(input))
			buffer = append(buffer[shift:],input[:shift]...)
			input = input[shift:]
		}

		front = buffer[:FRONT_BUFFER_SIZE]
		back = buffer[FRONT_BUFFER_SIZE:]

		fmt.Println("match:", match, length)
		fmt.Println("buffer:", string(front),string('|'), string(back))
		fmt.Println("input:", string(input))
		fmt.Println("output:", string(output))

	}
	fmt.Println()
	fmt.Println("Final Encoding:", string(output))
}

// aababacbaacbaad  -  aaaa  			-  a
// aababacbaacbaad	-	 aaaa aaba	-  02b
// bacbaacbaad			-	 aaab abac  -  23c
// baacbaad					-  abac baac  -  12a
// baad							-	 cbaa cbaa	-  03a
// d								-  cbaa d			-  00d

// encoded string  	- a22b23c12a03a30d
//										a02b23c12a03a00d

func findMatch(buffer []rune) (string, int){

	front := buffer[:window_length/2]
	back  := buffer[window_length/2:]
	//fmt.Println("front buffer:",string(front))
	//fmt.Println("back buffer:",string(back))
	length := 0
  offset := 0
	match := ""

	for i := range(front){
		//fmt.Println("index:", i,"character:", string(buffer[i]))
		if buffer[i] == back[0]{
			match=string(buffer[i])
			for j := i+1; j-i < len(back); j++{
				//fmt.Println(i, (j-i),j)
				//fmt.Println(string(back[j-i]), string(buffer[j]))
				if(buffer[j] == back[j-i] && j-i != len(back)-1){
					match+= string(buffer[j])
				}else{
					//fmt.Println("matched to:", match)
					if j - i > length {
						length = j - i
						offset = i
					}
					break
				}
			}
		}
	}

	//fmt.Println("longest match:", longest, "index of match:", index)
	match = string(buffer[offset:offset+length])
	test := ""
	test += strconv.Itoa(offset) + strconv.Itoa(length) + string(back[length])
	fmt.Println("match:", match, "encoded as:", test)
	return test, length
}

func Encode_Bitstream(encoding EncodingType , input string) []byte {

	//encode string into bitstream


	return []byte{}
}

func Decode_LZ(encoding EncodingType){

}