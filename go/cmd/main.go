package main

import (
	"fmt"
	"rosetta-code/pkg/linked"
)

func main() {
	n7 := linked.Node{Value: 7, Next: nil}
	n6 := linked.Node{Value: 6, Next: &n7}
	n5 := linked.Node{Value: 5, Next: &n6}
	n4 := linked.Node{Value: 4, Next: &n5}
	n3 := linked.Node{Value: 3, Next: &n4}
	n2 := linked.Node{Value: 2, Next: &n3}
	n1 := linked.Node{Value: 1, Next: &n2}
	l := linked.List{Head: &n1}

	fmt.Println(linked.KthToLastElement(l, 3))
}
