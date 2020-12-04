package main

import "fmt"

func main() {
	data := adding()

	first, second := divide(data, 3)
	fmt.Println("First : ", []string(first))
	second.print()
	fmt.Println("Word type to single string - ", first.toString())
	fmt.Println("Word type to byte - ", []byte(first.toString()))
	second.saveToFile("outputmain1", 0644)
	fmt.Println(readFile("outputmain1"))
}

func slice() string {
	return "This is second"
}
