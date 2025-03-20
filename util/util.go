package util

import (
	"container/heap"
	"fmt"
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
	remainder := len(binaryString)%8
	suffix := ""
	if remainder != 0 {
		fmt.Printf("additional bits: %d\n",remainder)
		for i := 0; i < 13 - remainder; i++ {
			suffix+="0"
		}
		fmt.Printf("suffix: %s\n", suffix)
		binaryCount := strconv.FormatInt(int64(remainder), 2)
		paddedBinaryCount := fmt.Sprintf("%03s", binaryCount)
		fmt.Printf("padded: %s\n", paddedBinaryCount)
		binaryString+=suffix+paddedBinaryCount;
	}
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
