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

func ReverseString(input string) string{
	
	fmt.Printf("The input string is : %s", input)
	//convert String to array of rune [97 90 81]
	convertStringToRune := []rune(input)
	
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
func ReverseWithoutSpace(input string){

	convertStringToRune := []rune(input)

	//loop and swap first and last elements
	for i,j:= 0,len(input)-1; i<len(input)/2; i,j = i+1, j-1{
		convertStringToRune[i],convertStringToRune[j] = convertStringToRune[j],convertStringToRune[i]
	}
	fmt.Printf("The Reversed String without using extra space is : %s \n", string(convertStringToRune))
}

/*
 	Given a positive integer, return its corresponding column title as appear in an Excel sheet.
 	1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 

*/
func ConvertToTitle(n int) string {
    finalString :=""
    for n>0{
        n--
        finalString += string(rune('A') + rune(n%26))
        n = n/26
    }
    return ReverseString(finalString)
}