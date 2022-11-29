package main

import (
	"log"
)

var number1 = 10

func main() {

	number2 := 22
	log.Println("Global variable , value of number1 : ", number1)
	log.Println("Local variable , value of number2 : ", number2)
	getValue()
}

func getValue() {
	number3 := 30

	log.Println("Global variable , value of number1 : ", number1)
	log.Println("Local variable , value of number3 : ", number3)

}
