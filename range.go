package main

import (
	"log"
)

func main() {

	var a = []uint{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for index, value := range a {
		log.Println("At index : ", index, "Value is : ", value)
	}
}
