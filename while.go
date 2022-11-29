package main

import (
	"log"
)

func main() {
	// here for loop works like while loop
	counter := 0
	
	for counter <= 5 {
		log.Println(counter)
		counter++
	}
}
