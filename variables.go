package main

import (
	"log"
)

func main() {
	var number1, number2, number3, number4, number5 float64

	number1 = 30
	number2 = 32.67
	number3 = 44
	number4 = 53
	number5 = 60.66

	average := (number1 + number2 + number3 + number4 + number5) / 5
	log.Println("Average is : ", average)

}
