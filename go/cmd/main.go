package main

import (
	"fmt"
	"rosetta-code/pkg/linked"
)

func main() {
	n5 := linked.Node{Value: 5, Next: nil}
	n4 := linked.Node{Value: 4, Next: &n5}
	n3 := linked.Node{Value: 3, Next: &n4}
	n2 := linked.Node{Value: 2, Next: &n3}
	n1 := linked.Node{Value: 1, Next: &n2}
	l := linked.List{Head: &n1}

	n5.Next = &n2

	fmt.Println(linked.FindLoopStart(l))
	fmt.Println(linked.FindLoopStartHard(l))
}
