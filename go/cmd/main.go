package main

import (
	"fmt"
	"rosetta-code/pkg/trees"
)

func main() {
	n1 := &trees.TreeNodeWithParent{Data: 4}
	n2 := &trees.TreeNodeWithParent{Data: -1}
	n3 := &trees.TreeNodeWithParent{Data: 3}
	n4 := &trees.TreeNodeWithParent{Data: 2}
	n5 := &trees.TreeNodeWithParent{Data: 1}

	n1.Left = n2
	n2.Parent = n1
	n2.Right = n3
	n3.Parent = n2
	n3.Left = n4
	n4.Parent = n3
	n4.Left = n5
	n5.Parent = n4

	fmt.Println(trees.FindNext(n1).ToString())
	fmt.Println(trees.FindNext(n2).ToString())
	fmt.Println(trees.FindNext(n3).ToString())
	fmt.Println(trees.FindNext(n4).ToString())
	fmt.Println(trees.FindNext(n5).ToString())
}
