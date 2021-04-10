package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

/* split into two slice and return two values */
func divide(w word, divide int) (word, word) {
	return w[:divide], w[divide:]
}

func (w word) toString() string {
	return strings.Join(w, ",")
}

func (w word) saveToFile(outputmain string, permission os.FileMode) error {
	return ioutil.WriteFile(outputmain, []byte(w.toString()), permission)
}

func from(filename string) ([]byte, error) {
	bs, error := ioutil.ReadFile(filename)
	return bs, error
}
