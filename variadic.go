package main

import (
	"log"
)

func main() {
	students := []string{"Ram", "Shyam", "Ganesh", "Sachin", "Virat", "Sourabh"}

	display(students...)
}

func display(students ...string) {
	for _, value := range students {
		log.Println(value)
	}
}
