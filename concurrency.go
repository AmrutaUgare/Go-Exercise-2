package main

import (
	"log"
	"runtime"
)

func main() {

	log.Println("OS : \t\t", runtime.GOOS)
	log.Println("Architecture : \t", runtime.GOARCH)
	log.Println("GoRoutines : \t", runtime.NumGoroutine())
	log.Println("CPU : \t\t", runtime.NumCPU())

	go first()
	second()

	log.Println("GoRoutines : \t", runtime.NumGoroutine())
	log.Println("CPU : \t\t", runtime.NumCPU())
}
func first() {

	log.Println("In First Function ...............")
}
func second() {

	log.Println("In Second Function ...............")

}
