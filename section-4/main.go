package main

import "fmt"

type person struct {
	firstName string
	lastName  string
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

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)
}
