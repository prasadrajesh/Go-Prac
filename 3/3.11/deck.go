package main

import "fmt"

type word []string

func (w word) print() {
	for i := range w {
		fmt.Println(w[i])
	}
}

func adding() word {
	data := word{}
	first := []string{"one", "two", "three"}
	second := []string{"four", "five", "six"}

	for _, first := range first {
		for _, second := range second {
			data = append(data, first+""+second)
		}
	}
	return data
}
