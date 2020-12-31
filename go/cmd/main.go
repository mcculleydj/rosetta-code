package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	n1 := &trees.TreeNode{Data: 1}
	n2 := &trees.TreeNode{Data: 2}
	n3 := &trees.TreeNode{Data: 3}
	n4 := &trees.TreeNode{Data: 4}
	n5 := &trees.TreeNode{Data: 5}
	// n6 := &trees.TreeNode{Data: 6}
	// n7 := &trees.TreeNode{Data: 7}

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n4.Left = n5
	// n3.Right = n7

	fmt.Println(trees.IsBalancedWrapper(n1))
}
