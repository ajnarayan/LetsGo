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
A more general version of Push(). Given a list, an index 'n' in the range 0..length,and a data element, 
add a new node to the list so that it has the given index.
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

/*
given a list that is sorted in increasing order, and a single node, 
inserts the node into the correct sorted position in the list
*/
func SortedInsert(newNode *Node, headref **Node){
	head := *headref
	if head != nil && newNode !=nil{
		current := head
		prev := current
		if head.Length() == 1{
			if head.Data > newNode.Data{
				Push(headref, newNode.Data)
			}else{
				
			}
		} 
		for i:=0; i<head.Length(); i++ {
			if newNode.Data>prev.Data && newNode.Data < current.Data{
				break
			}
			prev = current
			current = current.Next
		}
		prev.Next = newNode
		newNode.Next = current
	}
}