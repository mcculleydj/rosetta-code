package main

import (
	"fmt"
	"rosetta-code/pkg/iterables"
)

func main() {
	fmt.Println(iterables.OneAway("foo", "foo0"))
	fmt.Println(iterables.OneAway("foo", "fo"))
	fmt.Println(iterables.OneAway("fo0", "foo"))
	fmt.Println(iterables.OneAway("foo", "bar"))
	fmt.Println(iterables.OneAway("foo", "foooo"))
}
