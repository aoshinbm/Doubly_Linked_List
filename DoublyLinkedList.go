package main

import "fmt"

type Node struct {
	data int
	Next *Node
	Prev *Node
}
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dl *DoublyLinkedList) Display() {
	if dl.Head == nil {
		fmt.Println("Doubly Linked List is empty.")
		return
	}
	current := dl.Head
	fmt.Println("Doubly Linked List values:")
	for current != nil {
		fmt.Print("-->", current.data)
		current = current.Next
	}

}
func (dl *DoublyLinkedList) Append(value int) {
	newNode := &Node{data: value, Next: dl.Head}
	dl.Head = newNode
}
func main() {
	doubleLL := &DoublyLinkedList{}
	doubleLL.Append(1)
	doubleLL.Append(2)
	doubleLL.Append(3)
	doubleLL.Display()
}
