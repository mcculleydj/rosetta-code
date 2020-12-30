package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	ns := []int{1, 2, 3, 4, 7}
	fmt.Println(trees.MinBST(ns, nil).ToString())
}
