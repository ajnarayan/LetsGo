package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func PrintLL(ll *LinkedList) {
	temp := ll.Head
	if temp != nil {
		if temp.Next == nil {
			fmt.Printf("[%d,%p] --> ", temp.Data, temp.Next)
		} else {
			for temp != nil {
				fmt.Printf("[%d,%p] --> ", temp.Data, temp.Next)
				temp = temp.Next
			}
		}
		fmt.Printf("nil")
	}
}

func Append(ll *LinkedList, data int) {
	newNode := Node{data, nil}
	if ll.Head == nil {
		//This is first node to be inserted
		ll.Head = &newNode
	}
	if ll.Head.Next == nil {
		ll.Head.Next = &newNode
	} else {
		temp := ll.Head
		for temp != nil {
			temp = temp.Next
		}
		temp.Next = &newNode
	}
}

func Remove(ll *LinkedList, key int) {
	temp := ll.Head
	prev := ll.Head
	if ll.Head.Data == key {
		ll.Head = nil
	} else {
		for temp != nil {
			if temp.Data == key {
				break
			}
			prev = temp
			temp = temp.Next
		}
		if temp == nil {
			fmt.Println("Key not found in the list")
		} else {
			prev.Next = temp.Next
		}
	}
}

func Length(ll *LinkedList) int {
	length := 0
	temp := ll.Head
	for temp != nil {
		length++
		temp = temp.Next
	}
	return length
}

func Get(ll *LinkedList, key int) (bool, int) {
	temp := ll.Head
	isFound := false
	position := 0
	if ll.Head == nil {
		return isFound, position
	} else {
		for temp != nil {
			position++
			if temp.Data == key {
				isFound = true
				return isFound, position
			}
			temp = temp.Next
		}
	}
	return isFound, position
}
