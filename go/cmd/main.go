package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	n5 := &trees.TreeNode{Data: 5}
	n4 := &trees.TreeNode{Data: 4}
	n3 := &trees.TreeNode{Data: 3}
	n2 := &trees.TreeNode{Data: 2}
	n1 := &trees.TreeNode{Data: 1}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	fmt.Println(trees.BSTSequences(n1))
}
