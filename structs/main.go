package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "abc@cc",
			zipCode: "123",
		},
	}

	alex.updateName("Jimmy")
	fmt.Printf("%+v", alex)
}

func (pointToPerson *person) updateName(firstname string) {
	(*pointToPerson).firstname = firstname
}
