package main

import (
	"fmt"

)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int){
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node

}

func main() {
	var llist LinkedList
	llist = LinkedList{}
	llist.AddToHead(1)
	llist.AddToHead(3)

	fmt.Println(llist.headNode.property)

}
