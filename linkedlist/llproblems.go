package linkedlist
/*
The following problems are from :http://cslibrary.stanford.edu/105/LinkedListProblems.pdf
*/



// import (
// 	"fmt"
// )



/*
Prints the number of times given int occurs in a list
*/
func Count(head *Node, value int) int{
	count := 0
	if head != nil{
		temp := head
		for temp != nil{
			if temp.Data == value{
				count++
			}
			temp = temp.Next
		}
	}
	return count
}

/*
Returns the data value stored in nth position
*/
func GetNth(head *Node, n int) int{
	index := 0
	if head != nil{
		temp := head
		for temp != nil && index < head.Length() {
			if index == n{
				return temp.Data
			}
			temp = temp.Next
			index ++
		}
	}
	return 999
}

/*
 deallocates all of its memory and sets its head pointer to NULL (the empty list)
*/
func DeleteList(head **Node) {
	if *head != nil{
		current := *head
		for current != nil{
			next := current.Next
			current = nil
			current = next
		}
	*head = nil
	}
}
/*
 Pop() takes a non-empty list, deletes the head node, and returns the head node's data
*/
func Pop(head **Node) int{
	returnval := 0
	if *head == nil{
		return returnval
	}else{
		temp := *head
		returnval = temp.Data
		*head = temp.Next
		temp = nil
	}
	return returnval
}
/*
insert a new node at any index within a list.
*/
func InsertNth(index int, data int, headref **Node){
	head := *headref
	sizeoflist := head.Length()
	if index >=0 && index <= sizeoflist + 1{
		if index == 0{
			//append to head
			Push(headref, data)
		}else{
			current := head 
			prev := current
			for i :=0; i< sizeoflist ; i++{
				if i==index{
					break
				}
				prev = current
				current = current.Next					
			}
			newNode := Node{data, current}
			prev.Next = &newNode
		}
	}
}