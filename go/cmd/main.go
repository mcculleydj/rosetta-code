package main

import (
	"fmt"
	"rosetta-code/pkg/stacks"
)

func main() {
	q := stacks.Queue{}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	fmt.Println(q.ToString())

	fmt.Println(q.Remove())
	fmt.Println(q.ToString())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())

	fmt.Println(q.ToString())

}
