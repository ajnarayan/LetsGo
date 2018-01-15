package stringhelper

import "fmt"
import "reflect"

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