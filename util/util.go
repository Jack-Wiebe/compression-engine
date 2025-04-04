package util

import (
	"container/heap"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

//PriorityQueue
type Node struct {
	Value    string
	Priority int
	Index    int
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

// Len returns the number of items in the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less defines the order of items (min-heap: lower priority comes first)
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps two items in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority (lowest value in min-heap)
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // Mark as removed
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the queue
func (pq *PriorityQueue) Update(item *Node, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// Print the queue
func (pq PriorityQueue) Print(){
	temp := make(PriorityQueue, len(pq))
	copy(temp, pq)
	heap.Init(&temp) // Ensure the copy satisfies the heap property

	for temp.Len() > 0 {
		item := heap.Pop(&temp).(*Node)
		fmt.Printf("%s: %d\n", item.Value, item.Priority)
	}
	fmt.Println()
}

// Insert a value into the binary tree
func (n *Node) Insert(node Node) {
	if n == nil {
		return
	}

	// If the value is less than the current node's value, go left
	if node.Value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: node.Value, Priority: node.Priority}
		} else {
			n.Left.Insert(node)
		}
	} else { // Otherwise, go right
		if n.Right == nil {
			n.Right = &Node{Value:  node.Value, Priority: node.Priority}
		} else {
			n.Right.Insert(node)
		}
	}
}

// Pre-order traversal (root -> left -> right)
func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Printf("%s: %d \n", n.Value, n.Priority)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

// Post-order traversal (left -> right -> root)
func (n *Node) PostOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal()
	n.Right.PostOrderTraversal()
	fmt.Printf("%s ", n.Value)
}

// In-order traversal (left -> root -> right)
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Printf("%s ", n.Value)
	n.Right.InOrderTraversal()
}

//General
func PrintMap(input map[byte] int) {
	for char, freq := range input {
		fmt.Printf("%c: %d \n", char, freq)
	}
}

func FindKeyByValue(myMap map[string]string, targetValue string) (string, bool) {
	for key, value := range myMap {
		if value == targetValue {
			return key, true // Return the key and true if the value is found
		}
	}
	return "", false // Return an empty string and false if the value is not found
}

func ConvertToBinaryString(data []byte) string {
	// Convert each byte to its binary representation and join them
	var binaryStrings []string
	for _, b := range data {
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", b))
	}
	return strings.Join(binaryStrings, "")
}

func ConvertToByteArray(binaryString string) ([]byte, error){
	fmt.Println(binaryString)
	byteArray := make([]byte, 0, len(binaryString)/8)

	for i := 0; i < len(binaryString); i += 8 {
		chunk := binaryString[i : i+8]

		value, err := strconv.ParseUint(chunk, 2, 8)
		if err != nil {
			return nil, fmt.Errorf("invalid binary string: %v", err)
		}
		byteArray = append(byteArray, byte(value))
	}

	return byteArray, nil
}

func SaveTestImg() {
	// Create a 10x10 image

	pixel_array := [][] int {
		{1, 6, 4, 6, 2, 6, 4, 6, 1, 6, 4, 6}, // Row 0 (y=0)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, // Row 1 (y=10)
		{5, 6, 5, 6, 5, 6, 5, 6, 5, 6, 5, 6}, // Row 2 (y=20)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, // Row 3 (y=30)
		{3, 6, 4, 6, 3, 6, 4, 6, 3, 6, 4, 6}, // Row 4 (y=40)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, // Row 5 (y=50)
		{5, 6, 5, 6, 5, 6, 5, 6, 5, 6, 5, 6}, // Row 6 (y=60)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, // Row 7 (y=70)
		{1, 6, 4, 6, 2, 6, 4, 6, 1, 6, 4, 6}, // Row 8 (y=80)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, // Row 9 (y=90)
		{5, 6, 5, 6, 5, 6, 5, 6, 5, 6, 5, 6}, // Row 10 (y=100)
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},  // Row 11 (y=110)
	};

	color_array := []Color {
		{r:  0, g:  0, b:   0, a:   0},
		{r:170, g:170, b: 204, a: 255}, //rgb(170,170,204)
		{r:255, g:  0, b:   0, a: 255},	//rgb(255,0,0)
		{r:  0, g:128, b:   0, a: 255},	//rgb(0,128,0)
		{r:  0, g:  0, b: 255, a: 255},	//rgb(0,0,255)
		{r:170, g:153, b: 119, a: 255}, //rgb(170,153,119)
		{r:119, g:170, b: 119, a: 255},	//rgb(119,170,119)
		{r:255, g:255, b:   0, a: 255},	//rgb(255,255,0)
	}

	width, height := len(pixel_array[0]), len(pixel_array)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set all pixels to red
	//red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	for x := width-1; x >=0 ; x--{
		for y := 0; y < height; y++ {
			fmt.Println(x,y)
			fmt.Println(pixel_array[x][y])
			c := color_array[pixel_array[y][x]]
			img.Set(x, y, color.RGBA{R: c.r, G: c.g, B: c.b, A: 255})
		}
	}

	// Create the output file
	file, err := os.Create("test_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode as PNG and save
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func LoadImage(path string) ([][]Color, error){
	// Open image file
	file, err := os.Open(path) // or "input.jpg"
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil,err
	}
	defer file.Close()

	// Decode image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return  nil,err
	}
	fmt.Printf("Image format: %s\n", format)

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	//make empty 2D array for image
	image := make([][]Color, width)
	for i := range image {
		image[i] =  make([]Color, height)
	}

	// Read pixels
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get RGBA values (each range 0-65535)
			r, g, b, a := img.At(x, y).RGBA()

			// Convert to 8-bit values (0-255)
			var color Color
			color.r = uint8(r >> 8)
			color.g = uint8(g >> 8)
			color.b = uint8(b >> 8)
			color.a = uint8(a >> 8)
			image[x][y] = color


			fmt.Printf("Pixel at (%d,%d): R=%d, G=%d, B=%d, A=%d\n",
				x, y, color.r, color.g, color.b, color.a)
		}
	}

	return image, nil;
}

type Pixel struct {
	X int
	Y int
}

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

