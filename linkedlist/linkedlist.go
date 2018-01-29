package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func PrintLL(ll *List) {
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

func Append(ll *List, data int) {
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

func Remove(ll *List, key int) {
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

/*
(ll *List) - receiver (appears between func and nameoffunc())
ll is the receiver.
Same as classes, we can use ll.legth() to access this function
*/
func (ll *List) Length() int {
	length := 0
	temp := ll.Head
	for temp != nil {
		length++
		temp = temp.Next
	}
	return length
}

func Get(ll *List, key int) (bool, int) {
	temp := head(ll)
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

func head(ll *List) *Node {
	return ll.Head
}

/*
Adds new node at the head of the list
*/
func Push(ll *List, data int) {
	node := Node{data, nil}
	node.Next = ll.Head
	ll.Head = &node
}

/*
Building a linkedlist {5,4,3,2,1} using push and tail creation
*/
func BuildWithSpecialCase() *List {
	head := Node{1, nil}
	llHead := List{&head}
	for i := 2; i < 6; i++ {
		Push(&llHead, i)
	}
	return &llHead
}
