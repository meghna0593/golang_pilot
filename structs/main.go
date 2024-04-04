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
	contact   contactInfo
}

func main() {
	// person1 := person{"Jane", "Doe"} //one way of defining; not recommended
	// person1 := person{firstName: "Jane", lastName: "Doe"} // second way

	var person1 person // third way; defaulted with zero value

	person1.firstName = "Jane"
	person1.lastName = "Doe"

	fmt.Println(person1)
	fmt.Printf("%+v", person1)

}
