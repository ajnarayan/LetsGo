package main

/*

bool

string

int 	int8 	int16 	int32 	int64
uint 	unit8	unit16	uint32	uint64

byte 	//alias for uint8

rune 	//alias for int32 -- unicode code point

float32 	float64

complex64	complex128
*/
import (
	linkedlist "LetsGo/linkedlist"
	stringhelper "LetsGo/strings"
	"fmt"
)

func StringReverse() {
	var input = `!oG, olleH`
	fmt.Printf("\n----------------------Starting String Test-------------------------------------------\n")

	fmt.Println("String details:")

	stringhelper.PrintStringDetails(input)

	stringhelper.Reverse("!oG ,olleH")

	stringhelper.ReverseWithoutSpace(input)

	fmt.Printf("\n----------------------End of String Test-------------------------------------------\n")

}

func CaesarCipher() {
	var input = `This is a test`
	var shift = 2
	fmt.Printf("\n----------------------Starting CaesarCipher Test-------------------------------------------\n")
	fmt.Printf("Encrypting the characters \n")
	var encrypted = stringhelper.EncodeCharacters(input, shift)
	fmt.Printf("de-crypting the characters \n")
	shift = 2
	stringhelper.DecodeCharacters(encrypted, shift)
	fmt.Printf("\n----------------------End of CaesarCipher Test-------------------------------------------\n")
}

func LinkedList() {
	node := linkedlist.Node{1, nil}
	head := linkedlist.List{&node}
	fmt.Printf("\n----------------------Starting LinkedList Test-------------------------------------------\n")
	//fmt.Printf("Head Node data --> (%d,%p)", head.Head.Data, head.Head.Next)
	//linkedlist.PrintLL(&head)
	linkedlist.Append(&head, 2)
	//fmt.Printf("Head Node data --> (%d,%p)", head.Head.Data, head.Head.Next)
	linkedlist.PrintLL(&head)
	fmt.Println()
	linkedlist.Remove(&head, 2)
	linkedlist.PrintLL(&head)
	linkedlist.Append(&head, 4)
	length := head.Length()
	fmt.Printf("\n Length of the list is : %d", length)
	isFound, position := linkedlist.Get(&head, 1)
	if isFound {
		fmt.Printf("\n The element %d is found at %d position", 1, position)
	}
	var noofelements = 5
	fmt.Printf("\n LinkedList Inserting %d elements: \n", noofelements)
	newhead := linkedlist.AddToHead(noofelements)
	linkedlist.PrintLL(&linkedlist.List{newhead})
	fmt.Printf("\n LinkedList with special case using tail: \n")
	newtailhead := linkedlist.BuildWithSpecialCase()
	linkedlist.PrintLL(&linkedlist.List{newtailhead})
	fmt.Printf("\n LinkedList with special case using LocalRef: \n")
	newlocalhead := linkedlist.BuildWithSpecialCase()
	linkedlist.PrintLL(&linkedlist.List{newlocalhead})
	fmt.Printf("\n LinkedList CopyLinkedList: \n")
	copiedhead := linkedlist.CopyList(newlocalhead)
	linkedlist.PrintLL(&linkedlist.List{copiedhead})
	fmt.Printf("\n----------------------End of LinkedList Test-------------------------------------------\n")
}

func LinkedListProblems(){
	node := linkedlist.BuildWithSpecialCase()
	linkedlist.PrintLL(&linkedlist.List{node})
	var value = 2
	count := linkedlist.Count(node, value)
	fmt.Printf("\n Number of times %d occurs is %d \n",value, count)
	getNth := linkedlist.GetNth(node, value)
	fmt.Printf("\n Element in position %d is %d \n",value, getNth)
}
func main() {
	//StringReverse()
	//CaesarCipher()
	//LinkedList()
	LinkedListProblems()
}
