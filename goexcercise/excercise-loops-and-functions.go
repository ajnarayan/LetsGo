package main

/*
given a number x, we want to find the number z for which z² is most nearly x. (squareroot)

Computers typically compute the square root of x using a loop.
Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

Implement this in the func Sqrt provided.
A decent starting guess for z is 1, no matter what the input.
To begin with, repeat the calculation 10 times and print each z along the way.
See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

change the loop condition to stop once the value has stopped changing (or only changes by a very small amount).
See if that's more or fewer than 10 iterations.
Try other initial guesses for z, like x, or x/2.
How close are your function's results to the math.Sqrt in the standard library?

 the z² − x above is how far away z² is from where it needs to be (x),
 and the division by 2z is the derivative of z²,
 to scale how much we adjust z by how quickly z² is changing.
 This general approach is called Newton's method.
 It works well for many functions but especially well for square root.

 https://en.wikipedia.org/wiki/Newton%27s_method
*/

import (
	"fmt"
	"math"
)

const delta = 0.0000001

func checkdelta(oldvalue, newvalue float64) bool {
	if newvalue < 0.0 {
		newvalue = -newvalue
	}
	if math.Abs(newvalue-oldvalue) <= delta {
		return true
	}
	return false
}

func sqrt(x float64) (float64, int) {
	z := x / 2
	temp := 0.0
	i := int(0)
	for ; !checkdelta(temp, z); i++ {
		temp = z
		z -= (z*z - x) / (2 * z)
	}
	return z, i
}

func main() {
	var input = float64(6)
	result, iterations := sqrt(input)
	//fmt.Println(result == math.Sqrt(2))
	fmt.Printf("Expected value = %g, Actual value = %g, total iterations = %d ", result, math.Sqrt(input), iterations)
}
