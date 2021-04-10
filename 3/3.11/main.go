package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := adding()

	first, second := divide(data, 3)
	fmt.Println("First : ", []string(first))
	second.print()
	fmt.Println("Word type to single string - ", first.toString())
	fmt.Println("Word type to byte - ", []byte(first.toString()))
	second.saveToFile("outputmain1", 0644)
	by, err := from("outputmain11")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("data: ", by)
}

func slice() string {
	return "This is second"
}
