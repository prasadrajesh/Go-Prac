package main

import "fmt"

type contact struct {
	emailID string
	phoneNo string
}

type person struct {
	firstname string
	lastname  string
	contact   contact
}

func main() {
	// simple go struct
	// person1 := person{"rajesh", "prasad"}
	// fmt.Println(person1)

	// Another way to use struct
	person2 := person{firstname: "Rajesh", lastname: "Prasad"}
	fmt.Println(person2)

	// Null struct
	var person3 person
	fmt.Println(person3)
	//Field print
	fmt.Printf("%+v", person3)
	person3.firstname = "Raghav"
	person3.lastname = "Pal"
	fmt.Println("\n", person3)

	//Embedding struct
	person4 := person{
		firstname: "Vicky",
		lastname:  "Patel",
		contact: contact{
			emailID: "vickypatel@gmail.com",
			phoneNo: "800000000",
		},
	}

	fmt.Println(person4)
}
