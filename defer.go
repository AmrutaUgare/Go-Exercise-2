package main

import (
	"log"
)

func main() {

	// defer runs at last of the program
	defer log.Println("Hello")
	defer log.Println("!")
	defer log.Println("World")

}
