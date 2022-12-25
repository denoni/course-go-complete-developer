package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Alex",
		lastName:  "Gopher",
		contactInfo: contactInfo{
			email:   "alex@gopher.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Jim")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
