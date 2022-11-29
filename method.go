package main

import "log"

type student struct {
	name    string
	subject string
	marks   float32
}

func (s student) displayInformation() {

	log.Println("Name of student : ", s.name)
	log.Println("Subject of student : ", s.subject)
	log.Println("Marks of student : ", s.marks)
}

func main() {
	object := student{
		name:    "Sachin",
		subject: "Go",
		marks:   60,
	}

	object.displayInformation()
}
