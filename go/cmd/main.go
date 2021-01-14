package main

import (
	"fmt"
	"rosetta-code/pkg/files"
)

func main() {
	m, err := files.ReadFile("pkg/files/file_1.txt")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(files.MinWordDistance(m, "Hippolyte", "Lavater"))
}
