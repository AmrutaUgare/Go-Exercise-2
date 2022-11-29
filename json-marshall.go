// JSON Marshalling

package main

import (
	"encoding/json"
	"log"
)

type author struct {
	Name   string `json:"name"`
	Copies int    `json:"copies"`
	Age    int    `json:"age"`
}
type book struct {
	Id     int     `json : id`
	Price  float64 `json:"price"`
	Author author
}

func main() {

	string1 := author{Name: "Yashwant Kanetker", Copies: 100, Age: 50}

	string2 := book{Id: 101, Price: 200.0, Author: string1}

	byteArr, err := json.Marshal(string2)

	if err != nil {
		log.Println(err)
	}
	log.Println(string(byteArr))
}
