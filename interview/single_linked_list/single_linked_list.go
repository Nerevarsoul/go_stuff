package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type SingleLinkedList struct {
	head *Node
}

func (l *SingleLinkedList) len() int {
	lenght := 0
	node := l.head
	for node != nil {
		lenght++
		node = node.next
	}
	return lenght
}

func (l *SingleLinkedList) addNew(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *SingleLinkedList) printList() {
	node := l.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func (l *SingleLinkedList) reverseList() {
	node := l.head
	var prev_node *Node
	for node != nil {
		next_node := node.next
		node.next = prev_node
		prev_node = node
		node = next_node
	}
	l.head = prev_node
}

func main() {
	list := new(SingleLinkedList)
	a := new(Node)
	a.data = "Hello world!"
	b := new(Node)
	b.data = "Bye"
	list.addNew(a)
	list.addNew(b)
	list.reverseList()
	list.printList()
}
