package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

// Structs in Go is similar to a dictionary/class object(?) in Python
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// person1 := person{"Jane", "Doe"} //one way of defining; not recommended
	// person1 := person{firstName: "Jane", lastName: "Doe"} // second way

	// var person1 person // third way; defaulted with zero value

	// person1.firstName = "Jane"
	// person1.lastName = "Doe"

	person2 := person{
		firstName: "Michael",
		lastName:  "Scott",
		contactInfo: contactInfo{
			email:   "prisonmike@gmail.com",
			zipCode: "A0A0B0",
		},
	}

	// using pointers
	// person2Pointer := &person2
	// person2Pointer.updateFirstName("Prison Mike")
	// person2.print()

	// using pointers with lesser lines of code; shortcut
	person2.updateFirstName("Prison Mike")
	person2.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() { //receiver
	fmt.Printf("%+v", p)
}
