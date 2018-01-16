package main

import (
"fmt"
stringhelper "LetsGo/strings"
)


func StringReverse(){
	var input = `!oG, olleH`
	fmt.Printf("\n----------------------Starting String Test-------------------------------------------\n")
	
	fmt.Println("String details:")
	
	stringhelper.PrintStringDetails(input)
	
	stringhelper.Reverse("!oG ,olleH")
	
	stringhelper.ReverseWithoutSpace(input)

	fmt.Printf("\n----------------------End of String Test-------------------------------------------\n")

}

func CaesarCipher(){
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

func main(){
	//StringReverse()
	CaesarCipher()

	
}