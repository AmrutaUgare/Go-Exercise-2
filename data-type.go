package main

import (
	"log"
)

func main() {
	var variableUint uint8
	variableUint = 128

	log.Println("Value of unsigned variable is : ", variableUint)
	log.Printf("Data type of unsigned variable is : %T", variableUint)

	variableInt := -127
	log.Println("Value of int variable is : ", variableInt)
	log.Printf("Data type of int variable is : %T", variableInt)

	var variableFloat float32
	variableFloat = 200.0

	log.Println("Value of float variable is : ", variableFloat)
	log.Printf("Data type of float variable is : %T", variableFloat)

	var variableComplex complex64
	variableComplex = 12i + 23

	log.Println("Value of complex variable is : ", variableComplex)
	log.Printf("Data type of complex variable is : %T", variableComplex)

	var variableString string
	variableString = "Sachin"

	log.Println("Value of string variable is : ", variableString)
	log.Printf("Data type of string variable is : %T", variableString)

	var variableboolean bool
	variableboolean = true

	log.Println("Value of boolean variable is : ", variableboolean)
	log.Printf("Data type of boolean variable is : %T", variableboolean)

}
