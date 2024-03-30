package main

import (
	"fmt"
)

// Define a struct for the node
type Node struct {
	data interface{} // The data stored in the node (can be any type)
	next *Node       // Pointer to the next node
}

// Define a struct for the linked list
type SingleList struct {
	head *Node // Pointer to the first node in the list
}

// Implement the InsertFirst method on the SingleList struct
func (list *SingleList) InsertFirst(data interface{}) {
	newNode := Node{data: data, next: list.head} // Create a new node
	list.head = &newNode                         // Update the head pointer
}

// Function to create a new node (unchanged)
func newNode(data interface{}, next *Node) *Node {
	return &Node{data: data, next: next}
}

func (list *SingleList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	// Create a linked list
	myList := SingleList{}

	// Insert nodes into the list
	myList.InsertFirst(10)
	myList.InsertFirst(5)
	myList.InsertFirst(4)
	myList.InsertFirst(3)

	// Print the contents of the list
	myList.PrintList() // Output: 3 4 5 10
}
