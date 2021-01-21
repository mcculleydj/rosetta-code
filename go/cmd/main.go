package main

import (
	"fmt"
	"rosetta-code/pkg/bits"
)

func main() {
	// sum, err := bits.Add(-4, -10)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(sum)
	sum, err := bits.Add(9223372036854775807, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(sum)
}
