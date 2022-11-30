package main

import (
	"log"
)

func main() {

	num1 := 50
	num2 := 20

	summation(num1, num2)

	answer := subtractions(num1, num2)

	log.Println("Subtraction of ", num1, " - ", num2, " is ", answer)
}

func summation(num1, num2 int) {

	log.Println("Sum of ", num1, " + ", num2, " is ", num1+num2)

	result := summationByOtherWay(num1 + 2, num2 + 2)

	log.Println("In Another way, summation of num1 and num2 is : ", result)

}
func summationByOtherWay(num1, num2 int) int {
	return num1 + num2

}

func subtractions(num1, num2 int) int {
	return num1 - num2
}
