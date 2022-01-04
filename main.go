package main

import "fmt"

//The Node class has an integer typed variable with the name property.
//The class has another variable with the name nextNode, which is a node pointer
//Node class
type Node struct {
	property int
	nextNode *Node
}

//The LinkedList class has the headNode pointer as its property. By traversing
//to nextNode from headNode, you can iterate through the linked list
// LinkedList class
type LinkedList struct {
	headNode *Node
}

//Methods of the LinkedList class

//The AddToHead method adds the node to the start of the linked list.
//The AddToHead method of the LinkedList class has a parameter integer property.
//The property is used to initialize the node.
//A new node is instantiated and its property is set to the property parameter that's passed.
//The nextNode points to the current headNode of linkedList,
//and headNode is set to the pointer of the new node that's created

//AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = &node
}

//The IterateList method of the LinkedList class iterates from the headNode property
//and prints the property of the current head node.

//The iteration happens with the head node moves to nextNode of the headNode property
//until the current node is no longer equal to nil.

//IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

//The LastNode method of LinkedList returns the node at the end of the list. The list is
//traversed to check whether nextNode is nil from nextNode of headNode

//LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

//The AddToEnd method adds the node at the end of the list. In the following code,
//the current lastNode is found and its nextNode property is set as the added node

//AddToEnd method adds the node with property to the end
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

//The NodeWithValue method of LinkedList
//The list is traversed and checked to see whether the property value is equal
//to the parameter property

//NodeWithValue method returns Node given parameter property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

//AddAfter method adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {

	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	fmt.Println(linkedList.headNode.property)
	fmt.Println(linkedList.headNode.nextNode.property)
	fmt.Println("----")
	linkedList.IterateList()
	fmt.Println("----")
	fmt.Println(linkedList.LastNode().property)
	linkedList.AddToEnd(9)
	fmt.Println("----")
	linkedList.IterateList()
	fmt.Println("----")
	fmt.Println(linkedList.NodeWithValue(4).property)
	linkedList.AddAfter(4, 67)
	fmt.Println("----")
	linkedList.IterateList()

	// 1. REMOVE DUPLICATES

	//numeros := []int{2, 5, 6, 5, 1, 8, 5, 5, 1, 2, 4, 6, 2, 8, 5, 6, 2}

	// 2. RETURN ONLY THE REPEATED NAMES AND THEIR NUMBER OF OCCURRENCES.

}
