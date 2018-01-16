package strings

import "fmt"

func Reverse(input string){
	
	fmt.Printf("The input string is : %s", input)
	//convert String to array of rune [97 90 81]
	convertStringToRune := []rune(input)
	
	l := 0
	output := []rune(input)
	for i :=len(input)-1; i>=0 ; i--{
		output[l] = convertStringToRune[i]
		l++
	}

	fmt.Printf("The Reversed String is : %s \n", string(output))
}

/*
reverse string without using extra space
*/
func ReverseWithoutSpace(input string){

	convertStringToRune := []rune(input)

	//loop and swap first and last elements
	for i,j:= 0,len(input)-1; i<len(input)/2; i,j = i+1, j-1{
		convertStringToRune[i],convertStringToRune[j] = convertStringToRune[j],convertStringToRune[i]
	}
	fmt.Printf("The Reversed String without using extra space is : %s \n", string(convertStringToRune))
}