package main

import "fmt"

type word []string

func add() word {
	datas := word{}
	first := []string{"one", "two", "three"}
	second := []string{"four", "five", "six"}
	for _, firststr := range first {
		for _, secondstr := range second {
			datas = append(datas, firststr, secondstr)
		}
	}
	return datas
}

func (w word) print() {
	for i := range w {
		fmt.Println(w[i])
	}
}

func divide(w word) (word, word) {
	slice1 := word{}
	slice2 := word{}
	for i, fullslice := range w {
		if i%2 == 0 {
			slice1 = append(slice1, fullslice)
		}
		if i%2 != 0 {
			slice2 = append(slice2, fullslice)
		}
	}
	return slice1, slice2
}
