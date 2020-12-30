package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	root := new(trees.TreeNode)
	root.Data = 1
	root.Left = new(trees.TreeNode)
	root.Left.Data = 2
	root.Right = new(trees.TreeNode)
	root.Right.Data = 3
	root.Right.Right = new(trees.TreeNode)
	root.Right.Right.Data = 4
	root.Right.Right.Right = new(trees.TreeNode)
	root.Right.Right.Right.Data = 5

	q := new(trees.Queue)
	depths := []*trees.Queue{}
	ret := trees.ListDepths(root, 0, q, depths)

	for _, q := range ret {
		fmt.Println(q.ToString())
	}
}
