package main

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}
