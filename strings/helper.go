package strings

import "fmt"
import "reflect"

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


func PrintStringDetails(input string){
	//using goslices
	//output := make([]string, len(input))

	//indexing a string accesses individual bytes(%x), not characters.
	fmt.Printf("The type of the String[0] is : %s \n",reflect.TypeOf(input[0]))
	fmt.Printf("The location of the String is : %+s \n", &input)
	fmt.Printf("length of input String: %d\n",len(input))
	//indexing a string yields its bytes, not its characters: a string is just a bunch of bytes.
	//code point(rune) is an item represented by a single value (U+2318 , where 2318 is hexa representation of the char)
	for index, runeValue:= range input{
		fmt.Printf("The code point %#U starts at byte position %d\n", runeValue, index)		
	}
}