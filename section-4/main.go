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
	// Way one of declaring a struct
	// alex := person{"Alex", "Anderson"}
	// Way two of declaring a struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// Way three of declaring a struct. Assigns value of empty string ""
	// var alex person
	// alex.firstName="Alex"
	// alex.lastName="Anderson"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Bob",
		contactInfo: contactInfo{
			email:   "jim@bob.com",
			zipCode: 90210,
		},
	}

	jim.updateName("Bobby")

	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
