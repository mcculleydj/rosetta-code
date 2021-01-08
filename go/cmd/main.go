package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {

	n4 := &trees.BSTNode{Data: 4}

	t := &trees.BST{}
	t.Root = n4

	t.Insert(2)
	t.Insert(5)
	t.Insert(3)
	t.Insert(1)
	t.Insert(9)
	t.Insert(8)
	t.Insert(7)

	fmt.Println(t.ToString())
	t.Delete(8)
	fmt.Println(t.ToString())
	t.Delete(4)
	fmt.Println(t.ToString())
	t.Delete(4)
	fmt.Println(t.ToString())

	m := map[int]int{}
	for i := 0; i < 100000; i++ {
		m[t.Random().Data]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
