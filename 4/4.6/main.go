package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//Pointer for struct
	person1 := person{"Rajesh", "Prasad"}

	person1.adding("Lucky")
	person1.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) adding(firstname string) {
	(*p).firstName = firstname
}
