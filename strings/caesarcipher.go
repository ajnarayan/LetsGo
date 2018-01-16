package strings

import "fmt"
import "strings"

func EncodeCharacters(input string, shift int) string {
	output := []rune(strings.ToLower(input))
	fmt.Printf("The input string is %s\n", string(output))
	shiftchar := rune(shift)

	for i := range output {
		//fmt.Printf("%#U", output[i])
		if output[i] != rune(' ') {

			output[i] = output[i] + shiftchar
			if output[i] > rune('z') {
				output[i] = output[i] - rune(26)
			} else if output[i] < rune('a') {
				output[i] = output[i] + rune(26)
			}
		}

	}
	//fmt.Printf("%U", rune(26))
	//fmt.Printf("%U",rune('z'))
	fmt.Printf("The Encrypted String is : %s \n", string(output))
	return string(output)
}

func DecodeCharacters(input string, shift int) {
	output := []rune(strings.ToLower(input))
	shiftchar := rune(shift)
	for i := range output {
		if output[i] != rune(' ') {
			output[i] = output[i] - shiftchar
			if output[i] > rune('z') {
				output[i] = output[i] - rune(26)
			} else if output[i] < rune('a') {
				output[i] = output[i] + rune(26)
			}
		}
	}

	fmt.Printf("The Decrypted value is %s \n", string(output))
}
