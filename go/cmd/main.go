package main

import (
	"fmt"
	"rosetta-code/pkg/graphs"
)

func main() {
	n4 := graphs.Node{Data: 4, Adj: []*graphs.Node{}}
	n3 := graphs.Node{Data: 3, Adj: []*graphs.Node{&n4}}
	n2 := graphs.Node{Data: 2, Adj: []*graphs.Node{&n4}}
	n1 := graphs.Node{Data: 1, Adj: []*graphs.Node{&n2, &n3}}

	fmt.Println(graphs.PathExists(&n1, &n4))
	fmt.Println(graphs.PathExists(&n2, &n3))
}
