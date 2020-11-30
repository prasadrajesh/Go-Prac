// package main

// import "fmt"

// func main() {
// 	data := []string{slice(), "two", "three"}
// 	data = append(data, "four", "five")
// 	for i, j := range data {
// 		fmt.Println(i, j)
// 	}
// }

// func slice() string {
// 	return "This is second"
// }

package main

import "fmt"

func main() {
	data := []string{"one", "two"}
	data = append(data, "three", test())
	// for i := 0; i < 100; i++ {
	// 	fmt.Print(i)
	// 	fmt.Println("")
	// }
	for i, temp := range data {
		fmt.Println(i, temp)
	}

}

func test() string {
	return "four"
}
