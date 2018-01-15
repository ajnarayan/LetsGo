package main

import (
"fmt"
stringhelper "LetsGo/stringhelper"
)

// import "strings"
/*
	Go- slicing
	* Very similar to python slicing (except there is no negative index)
	A slice descriptor of an array has (pointer to array, length of segment and capacity of segment) 
	Go source code is always UTF-8.
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	Those sequences represent Unicode code points, called runes.
	No guarantee is made in Go that characters in strings are normalized.


	%x- hex byte
	%q - ascii 
	%s - string
	
	A string `⌘`(%s) is represented in hex bytes e2 8c 98 (%x), these bytes are UTF-8 encoding of 
	hexadecimal value 2318 (%q - \u2318)

	In go only String literals `` are in UTF-8
	A string can contain other arbitrary bytes ("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	Therefore, strings can contain arbitrary bytes, but when constructed from string literals, those bytes are (almost always) UTF-8.

	Using Runes: 
	eg1: 
	r := rune('a')
	fmt.Println(r, string(r)) //outputs 97 a

	eg2:
	for i, r := range "abc" {
    fmt.Printf("%d - %c (%v)\n", i, r, r)
	}
	outputs - 
	0 - a (97)
	1 - b (98)
	2 - c (99)

*/

func main(){
	var input = `!oG, olleH`
	
	fmt.Printf("Starting String Test \n")
	
	fmt.Println("String details:")
	
	stringhelper.PrintStringDetails(input)
	
	fmt.Printf("\nThe Reversed String is : %s \n", Reverse("!oG ,olleH"))
	
	fmt.Printf("\n The Reversed String without using extra space is : %s \n", ReverseWithoutSpace(input))
}


func Reverse(input string) string{
	//convert String to array of rune [97 90 81]
	convertStringToRune := []rune(input)

	fmt.Printf("The String is : %s \n", string(convertStringToRune))
	
	l := 0
	output := []rune(input)
	for i :=len(input)-1; i>=0 ; i--{
		output[l] = convertStringToRune[i]
		l++
	}

	return string(output)
}

/*
reverse string without using extra space
*/
func ReverseWithoutSpace(input string) string{

	convertStringToRune := []rune(input)

	//loop and swap first and last elements
	for i,j:= 0,len(input)-1; i<len(input)/2; i,j = i+1, j-1{
		convertStringToRune[i],convertStringToRune[j] = convertStringToRune[j],convertStringToRune[i]
	}
	return string(convertStringToRune)

}

