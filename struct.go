package main

import (
	"log"
)

type house struct {
	noRooms int32
	price   float64
	city    string
}

func main() {
	log.Println(house{noRooms: 2, price: 7000000.0, city: "Pune"})

	house1 := house{1, 4000000.0, "Pune"}
	log.Println(house1)

	house1.noRooms = 3
	house1.price = 10000000.0
	house1.city = "Pune"

	log.Println("Rooms : ", house1.noRooms, "Price : ", house1.price, "City : ", house1.city)
}
