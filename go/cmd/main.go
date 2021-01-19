package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	n1 := &trees.TreeNode{Data: 0}
	n2 := &trees.TreeNode{Data: 0}
	n3 := &trees.TreeNode{Data: 3}
	n4 := &trees.TreeNode{Data: 3}

	n1.Left = n2
	n2.Left = n3
	n2.Right = n4

	fmt.Println(trees.PathsToSum(n1, 3, 0, map[int]int{}))
}
