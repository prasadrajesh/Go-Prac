package main

import "fmt"

// Without pointer update slice
func main() {
	read := []string{"Hi", "there", "is", "a", "tree"}
	update(read)
	fmt.Println(read)
}

func update(update []string) {
	update[4] = "Animal"
}
