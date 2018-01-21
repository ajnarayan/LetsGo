package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

// func Add(node *LinkedList) (bool, LinkedList) {
// 	isSuccess := bool(false)
// 	if node != nil {
// 		Head.Next = node
// 		node.Next = nil
// 		isSuccess = true
// 	}
// 	return isSuccess, Head
// }

func PrintLL(ll *LinkedList) {
	temp := ll.Head
	if temp != nil {
		if temp.Next == nil {
			fmt.Printf("[%d,%p]", temp.Data, temp.Next)
		} else {
			for temp != nil {
				fmt.Printf("[%d,%p] --> ", temp.Data, temp.Next)
				temp = temp.Next
			}
			fmt.Printf("nil")
		}
	}
}

func Add(ll *LinkedList, data int) {
	newNode := Node{data, nil}
	if ll.Head == nil {
		//This is first node to be inserted
		ll.Head = &newNode
	}
	if ll.Head.Next == nil {
		ll.Head.Next = &newNode
	} else {
		temp := ll.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newNode
	}
}
