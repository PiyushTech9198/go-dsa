package main

import (
	"fmt"
)

// Node struct for the linked list
type Node struct {
	data interface{}
	next *Node
}

// SingleLinkedList struct implements the LinkedList interface
type SingleLinkedList struct {
	head *Node
}

// Implement InsertAtFirst for SingleLinkedList
func (sll *SingleLinkedList) InsertAtFirst(data interface{}) {
	newNode := &Node{data: data, next: sll.head}
	sll.head = newNode
}

// Implement InsertAtLast for SingleLinkedList
func (sll *SingleLinkedList) InsertAtLast(data interface{}) error {
	if sll.IsEmpty() {
		sll.head = &Node{data: data}
		return nil
	}

	newNode := &Node{data: data, next: nil}
	temp := sll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	return nil
}

// Implement IsEmpty for SingleLinkedList
func (sll *SingleLinkedList) IsEmpty() bool {
	return sll.head == nil // Corrected condition: nil head means empty list
}

// InsertAfter inserts a new node with data after a specific node
func (sll *SingleLinkedList) InsertAfter(insertAfterData interface{}, data interface{}) {
	temp := sll.head
	for temp != nil {
		if temp.data == insertAfterData {
			newNode := &Node{data: data, next: temp.next}
			temp.next = newNode
			return // Inserted, exit the loop
		}
		temp = temp.next
	}
	fmt.Println("InsertAfter failed: data not found") // Indicate if insertion failed
}

// DeleteFirst removes the first node from the list
func (sll *SingleLinkedList) DeleteFirst() error {
	if sll.IsEmpty() {
		return fmt.Errorf("DeleteFirst failed: list is empty")
	}

	// Update head to point to the next node
	sll.head = sll.head.next

	// Set the first node's next pointer to nil to avoid potential memory leaks
	temp := sll.head
	temp.next = nil // Previously this was: sll.head.next = temp (incorrect)

	return nil
}

// Search finds a node with the specified data
func (sll *SingleLinkedList) Search(data interface{}) *Node {
	temp := sll.head
	for temp != nil {
		if temp.data == data {
			return temp
		}
		temp = temp.next
	}
	return nil // Data not found
}

// DeleteAtNode removes a specific node from the list
func (sll *SingleLinkedList) DeleteAtNode(node *Node) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty")
		return
	}

	// Check if the node to be deleted is the head
	if sll.head == node {
		sll.head = sll.head.next
		return
	}

	// Traverse the list to find the node before the target node
	temp := sll.head
	for temp != nil && temp.next != nil {
		if temp.next == node {
			// Found the node before the target node
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}

	// If loop finishes without finding the node
	fmt.Println("Node not found in the list")
}

// DeleteLast removes the last node from the list
func (sll *SingleLinkedList) DeleteLast() {
	if sll.IsEmpty() {
		fmt.Println("List is empty cannot delete last element")
		return
	}

	var prev *Node
	node := sll.head
	for node.next != nil {
		prev = node
		node = node.next
	}

	if prev == nil { // Only one node in the list
		sll.head = nil
	} else {
		prev.next = nil
	}
}

// PrintList displays the data of each node in the list
func (sll *SingleLinkedList) PrintList() {
	temp := sll.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
	fmt.Println("---- End of List ----")
}

// Iterator for SingleLinkedList
type SLLIterator struct {
	current *Node
}

// Create a new iterator for a SingleLinkedList
func (sll *SingleLinkedList) Iterator() *SLLIterator {
	return &SLLIterator{current: sll.head}
}

// HasNext checks if there are more elements to iterate over
func (it *SLLIterator) HasNext() bool {
	return it.current != nil
}

// Next returns the data of the current node and advances the iterator
func (it *SLLIterator) Next() interface{} {
	if !it.HasNext() {
		fmt.Println("No more elements to iterate")
		return nil
	}
	data := it.current.data
	it.current = it.current.next
	return data
}

func main() {
	// Create a new empty SingleLinkedList
	sll := &SingleLinkedList{} // Initialize with nil head

	if !sll.IsEmpty() { // Check for empty list
		fmt.Println("List is empty")
	} else {
		fmt.Println("List is not empty") // This shouldn't print now
	}

	sll.InsertAtFirst(1)
	sll.InsertAtFirst(2)
	sll.InsertAtFirst(3)
	sll.InsertAtFirst(4)

	fmt.Println("Original List:")
	sll.PrintList()

	// Delete the node with data 4
	sll.DeleteAtNode(sll.Search(4))

	fmt.Println("List after deleting 4:")
	sll.PrintList()

	// Use the iterator
	iterator := sll.Iterator()
	fmt.Println("Iterating through the list:")
	for iterator.HasNext() {
		data := iterator.Next()
		fmt.Println(data) // Prints list values: 3 2 1
	}
}
