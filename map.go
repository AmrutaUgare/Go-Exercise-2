package main

import (
	"log"
)

func main() {

	map1 := map[string]int{}

	map1["Pune"] = 1
	map1["Mumbai"] = 2
	map1["Kolhapur"] = 3
	map1["Delhi"] = 4
	map1["Jaipur"] = 5

	log.Println(map1)

	// display value of perticula key
	log.Println("Value of map1[Delhi] is : ", map1["Delhi"])

	// add in map
	map1["Rajastan"] = 6
	map1["Banglore"] = 7

	log.Println(map1)

	// show map using range

	for index, value := range map1 {

		log.Println("Index : ", index, "||", "Value : ", value)
	}

	// delete

	delete(map1, "Mumbai")
	log.Println("After deletion ", map1)

}
