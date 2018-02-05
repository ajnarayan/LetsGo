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

// /*

// */
// func GetNth()