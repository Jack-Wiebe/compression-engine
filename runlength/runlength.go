package runlength

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Encode_Runlength(file *os.File) {
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking:", err)
		return
	}
	fmt.Println("Run Length Encoding", file.Name())
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	//
	//RUN LENGTH ENCODER
	//

	count := 0
	output := ""
	var char_pointer string

	for scanner.Scan() {
		char := scanner.Text()
		if char_pointer == ""{
			count++
			char_pointer = char
		}else if char_pointer != char {

			output+=strconv.Itoa(count)
			output+=","
			output+=char_pointer
			char_pointer = char;
			count = 1
		}else{
			count++
		}
	}

	output+=strconv.Itoa(count)
	output+=","
	output+=char_pointer

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println()
	fmt.Println(output)

	//output to binary file
	os.WriteFile("test/output.txt", []byte(output), 0644)

}

func Decode_Runlength(){

	data, err := os.ReadFile("test/output.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	text := string(data)
	fmt.Println()
	fmt.Println(text)

	count := 0
	str := ""
	output := ""
	flag := false

	fmt.Println()
	for _,char := range text{
		if flag{
			for i := 0; i < count; i++ {
				output+=string(char)
			}
			str=""
			flag = false
			continue
		}
		if char == ','{
			count,_ = strconv.Atoi(str)
			flag = true
			continue
		}

		str+=string(char)
	}

	fmt.Println(output)

}
