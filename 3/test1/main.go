package main

import "fmt"

// import "fmt"

// func main() {
// 	data := add()
// 	fmt.Println("name- ", data)
// }

// func add() string {
// 	return "rajesh"
// }

// func main() {
// 	datas := []string{"one", "two", "three"}
// 	for i, temp := range datas {
// 		fmt.Println("name- ", i, temp)
// 	}
// 	for _, temp := range datas {
// 		fmt.Println("name- ", temp)
// 	}
// }

/*type declaration */
// func main() {
// 	datas := word{"one", "two", "three"}
// 	for i, temp := range datas {
// 		fmt.Println("name- ", i, temp)
// 	}
// 	for _, temp := range datas {
// 		fmt.Println("name- ", temp)
// 	}
// }

func main() {
	datas := add()
	// fmt.Println(datas)
	// datas.print()
	temp, temp1 := divide(datas)
	fmt.Println(temp, temp1)
}
