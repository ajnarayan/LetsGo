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

func (head *Node) PrintLLHead(){
	temp:= head
	if temp != nil{
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

func (head *Node) Length() int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
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
 The parameter has the word
"ref" in it as a reminder that this is a "reference" (**Node) pointer to the
head pointer instead of an ordinary (*Node) copy of the head pointer.
*/
func Push(headref **Node, data int) {
	node := Node{data, nil}
	node.Next = *headref
	*headref = &node
}

/*
Building a linkedlist {1,2,3,4,5} using push and tail creation
Push(NIL, 2):
   node = [2,nil]
   node.next = *NIL
   NIL = &[2,nil]
*/
func BuildWithSpecialCase() *Node {
	//Pointer to first node
	head := &Node{1, nil}
	//tail points to first node
	tail := head

	//we pass tail.next to push
	for i := 2; i < 6; i++ {
		Push(&tail.Next, i)
		tail = tail.Next
	}
	fmt.Printf("Building LL with head=%d, Tail=%d and size=%d \n", head.Data, tail.Data, head.Length())
	return head
}

func BuildWithLocalRef() *Node {
	head := Node{1, nil}
	localref := &head
	for i := 1; i < 6; i++ {
		Push(&localref, i)
		localref = &*localref.Next
	}
	return &head
}

func BuildWithDummySpecialCase() *Node {
	dummy := Node{0, nil}
	tail := &dummy
	for i := 1; i < 6; i++ {
		Push(&tail.Next, i)
		tail = tail.Next
	}
	return dummy.Next
}

func AddToHead(noofitems int) *Node {
	head := &Node{1, nil}
	for i := 1; i < noofitems; i++ {
		Push(&head, i)
	}
	return head
}

func AppendNode(headref **Node, data int) *Node {
	current := *headref
	if current == nil {
		Push(headref, data)
	} else {
		for current.Next != nil {
			current = current.Next
		}
		Push(&current.Next, data)
	}
	return *headref
}

func CopyList(head *Node) *Node {
	current := head
	tail := head
	i := 0
	node := Node{head.Data, nil}
	for current != nil {
		if i == 0 {
			tail = &node
		} else {
			tail.Next = &Node{current.Data, nil}
			tail = tail.Next
		}
		current = current.Next
		i++
	}
	return &node
}
