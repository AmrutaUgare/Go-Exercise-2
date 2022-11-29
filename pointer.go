package main

import (
	"log"
)

func main() {

	a := 10

	log.Println("Value of a : ", a)
	log.Println("Address of a : ", &a)

	var b *int

	b = &a
	log.Println("Address of b : ", b)

	c := *b
	log.Println("Value of c : ", c)
}
