package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	headNode *Node
	length   int
}

func (l *LinkedList) AddToHead(value int) {
	var node = Node{}
	node.value = value
	if node.next != nil {
		node.next = l.headNode
	}
	l.headNode = &node

}

func (l *LinkedList) IterateList() {
	var node *Node

	for node = l.headNode; node != nil; node = node.next {
		fmt.Println(node.value)

	}

	/*
		for l.headNode != nil {
			fmt.Println(node.property)
			l.headNode = l.headNode.nextNode
		}
	*/
}

func main() {
	var llist LinkedList
	llist = LinkedList{}

	// HeadNode
	llist.AddToHead(1)
	llist.AddToHead(3)
	llist.AddToHead(5)

	fmt.Println(llist.headNode.value)

	// IterateList

	llist.IterateList()
}
