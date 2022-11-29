/* Recursion - Accept a no. from user and find its factorial */
package main

import (
	"fmt"
)

func main() {
	number := 5

	answer := factorial(number)
	fmt.Println("Factorial of ", number, " is : ", answer)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return (number * factorial(number-1))
	}
}
