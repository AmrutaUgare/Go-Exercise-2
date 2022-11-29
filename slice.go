package main

import (
	"log"
)

func main() {
	message := "Hello World"

	log.Println(message[0:5])
	log.Println(message[5:])

	firstWord := message[0:5]
	secondWord := message[5:]

	log.Println("String is : ", firstWord, secondWord)

}
