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
	n6 := &trees.TreeNode{Data: 6}
	n7 := &trees.TreeNode{Data: 7}
	n8 := &trees.TreeNode{Data: 8}
	n9 := &trees.TreeNode{Data: 9}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n4.Left = n6
	n4.Right = n7
	n7.Left = n8
	n7.Right = n9

	n10 := &trees.TreeNode{Data: 2}
	n11 := &trees.TreeNode{Data: 4}
	n12 := &trees.TreeNode{Data: 5}
	n13 := &trees.TreeNode{Data: 6}
	n14 := &trees.TreeNode{Data: 7}
	n15 := &trees.TreeNode{Data: 8}
	n16 := &trees.TreeNode{Data: 9}

	n10.Left = n11
	n10.Right = n12
	n11.Left = n13
	n11.Right = n14
	n14.Left = n15
	n14.Right = n16

	fmt.Println(trees.CheckSubtree(n1, n10))
}
