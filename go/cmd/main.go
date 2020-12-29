package main

import (
	"fmt"
	"rosetta-code/pkg/stacks"
)

func main() {
	ss := stacks.StackSet{StackCapacity: 5}
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)

	fmt.Println(ss.ToString())

	fmt.Println(ss.PopAt(1))
	fmt.Println(ss.PopAt(1))
	fmt.Println(ss.PopAt(1))
	fmt.Println(ss.PopAt(1))
	fmt.Println(ss.PopAt(1))

	fmt.Println(ss.ToString())

	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())

	fmt.Println(ss.ToString())

	fmt.Println(ss.Pop())
	fmt.Println(ss.ToString())
}
