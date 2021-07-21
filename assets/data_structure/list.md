# List | [Home](../../README.md)

A list is a sequence of elements. Each element can be connected to another with a link in a forward or backward direction. The element can have other payload properties. This data structure is a basic type of container. Lists have a variable length and developer can remove or add elements more easily than an array. Data items within a list need not be contiguous in memory or on disk. Linked lists were proposed by Allen Newell, Cliff Shaw, and Herbert A. Simon at RAND Corporation.

To get started, a list can be used in Go, as shown in the following example; elements are added through the PushBack method on the list, which is in the container/list package:

```
package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(21)
	intList.PushBack(31)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
```

Lists also be used as a base for other data structures, such as stack and queue. Lists can be used to store lists of users, car parts, ingredients, to-do items, and various other such elements. Lists are the most commonly used linear data structures. These were introduced in the lisp programming language

Let's discuss the operations related to add, update, remove, and lookup on linked list and doubly linked list in the following section.

## LinkedList

LinkedList is a sequence of nodes that have properties and a reference to the next node in the sequence. It is a linear data structure that is used to store data. The data structure permits the addition and deletion of components from any node next to another node. They are not stored contiguously in memory, which makes them different arrays.

## The Node class

The Node class has an integer typed variable with the name property. The class has another variable with the name nextNode, which is a node pointer. Linked list will have a set of nodes with integer properties, as follows:

```
//Node class
type Node struct {
    property int
    nextNode *Node
} 
```

## The LinkedList class

The LinkedList class has the headNode pointer as its property. By traversing to nextNode from headNode, you can iterate through the linked list, as shown in the following code:

```
// LinkedList class
type LinkedList struct {
    headNode *Node
}
 
```

The different methods of the LinkedList class, such as AddtoHead, IterateList, LastNode, AddtoEnd, NodeWithValue, AddAfter, and the main method, are discussed in the following sections.

## The AddToHead method

The AddToHead method adds the node to the start of the linked list. The AddToHead method of the LinkedList class has a parameter integer property. The property is used to initialize the node. A new node is instantiated and its property is set to the property parameter that's passed. The nextNode points to the current headNode of linkedList, and headNode is set to the pointer of the new node that's created, as shown in the following code:

```
//AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
    var node = Node{}
    node.property = property
    if node.nextNode != nil {
        node.nextNode = linkedList.headNode
    }
    linkedList.headNode = &node
}
```

