package main

import (
	"fmt"
)

// Define a struct for the node
type Node struct {
	data interface{} // The data stored in the node (can be any type)
	next *Node       // Pointer to the next node
}

// Function to create a new node
func newNode(data interface{}, next *Node) *Node {
	return &Node{data: data, next: next}
}

func main() {
	// Create some nodes
	node1 := newNode(10, nil)
	fmt.Println(node1)
}
