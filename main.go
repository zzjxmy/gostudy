package main

import (
	"fmt"
	user "mypro/Test"
	user1 "mypro/User"
)

func main() {
	var name string
	var nameA string = "坚果"
	nameB := "坚果"
	var nameC, ageC = "hello", 1
	nameD, ageD := "world", 2

	fmt.Println("%C", name)
	fmt.Println("%C", nameA)
	fmt.Println("%C", nameB)
	fmt.Println("%C", nameC)
	fmt.Println("%C", ageC)
	fmt.Println("%C", nameD)
	fmt.Println("%C", ageD)

	fmt.Println("Hello world")
	fmt.Println("坚果是多维度最帅的男人")
	user.H()
	user1.H()
}
