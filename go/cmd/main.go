package main

import (
	"fmt"
	"rosetta-code/pkg/linked"
)

func main() {
	n7 := linked.Node{Value: 1, Next: nil}
	n6 := linked.Node{Value: 2, Next: &n7}
	n5 := linked.Node{Value: 9, Next: &n6}
	n4 := linked.Node{Value: 5, Next: &n5}
	l1 := linked.List{Head: &n4}

	n3 := linked.Node{Value: 8, Next: nil}
	n2 := linked.Node{Value: 5, Next: &n3}
	n1 := linked.Node{Value: 3, Next: &n2}
	l2 := linked.List{Head: &n1}

	// sumList := linked.RightToLeftSum(l1, l2)
	sumList := linked.LeftToRightSum(l1, l2)

	fmt.Println(sumList.ToString())
}
