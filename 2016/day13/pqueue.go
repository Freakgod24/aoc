package main

import (
	"fmt"
)

type Node struct {
	object   interface{}
	priority uint
	previous *Node
	next     *Node
}

type PriorityQueue struct {
	root *Node
}

func (pq *PriorityQueue) push(object interface{}, priority uint) {
	// The list is empty
	if pq.root == nil {
		pq.root = &Node{
			object,
			priority,
			nil,
			nil,
		}
		return
	}

	if pq.root.priority >= priority {
		// we are at the root of the list
		newNode := Node{object, priority, nil, pq.root}
		pq.root.previous = &newNode
		pq.root = &newNode
		return
	}

	// We need to traverse the list
	currentNode := pq.root
	for currentNode.priority < priority {
		if currentNode.next != nil {
			currentNode = currentNode.next
			continue
		}

		// We are at the end of the list
		newNode := Node{object, priority, currentNode, nil}
		currentNode.next = &newNode
		currentNode = &newNode
		return
	}

	// We are in the middle of the list
	newNode := Node{object, priority, currentNode.previous, currentNode}
	currentNode.previous.next = &newNode
	currentNode.previous = &newNode
}

func (pq *PriorityQueue) pop() interface{} {
	node := pq.root

	// We have an empty queue
	if node == nil {
		return nil
	}

	// We have a single element queue
	if node.next == nil {
		pq.root = nil
		return node.object
	}

	// We have a many elements queue
	pq.root = pq.root.next
	pq.root.previous = nil
	return node.object
}

func (pq *PriorityQueue) contains(object interface{}) bool {
	// We have an empty queue
	if pq.root == nil {
		return false
	}

	// Root is the object
	if pq.root.object == object {
		return true
	}

	// Traverse the list to find object
	currentNode := pq.root
	for currentNode.object != object {
		// We are at the end of the list and the object was not found
		if currentNode.next == nil && currentNode.object != object {
			return false
		} else if currentNode.next != nil {
			currentNode = currentNode.next
		}
	}

	// The object was found
	return true
}

func (pq *PriorityQueue) print() {
	fmt.Println()
	fmt.Print("[")
	currentNode := pq.root
	for currentNode.next != nil {
		fmt.Print(currentNode.priority, ", ")
		currentNode = currentNode.next
	}
	fmt.Print(currentNode.priority, "]\n")
}
