package linkedlist

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