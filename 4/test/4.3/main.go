package main

import "fmt"

type contact struct {
	emailID string
}

type person struct {
	firstname string
	lastname  string
	contact   contact
}

func main() {
	// simplest method for struct
	person1 := person{"rajesh", "prasad", contact{emailID: "dd"}}
	fmt.Println(person1)

	// second method
	person2 := person{firstname: "vicky", lastname: "patel"}
	fmt.Println(person2)

	// null value
	var person3 person
	fmt.Println(person3)

	fmt.Printf("%+v", person3)
	person3.firstname = "karan"
	person3.lastname = "raval"
	fmt.Println(person3)

	person4 := person{
		firstname: "jyoti",
		lastname:  "",
		contact:   contact{emailID: "jyotiprasad@gmail.com"},
	}

	fmt.Println(person4)
}
