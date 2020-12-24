package main

import (
	"fmt"
	"rosetta-code/pkg/one"
)

func main() {
	fmt.Println(one.CheckPermutation("doof", "food"))
	fmt.Println(one.CheckPermutation("baby", "food"))
	fmt.Println(one.CheckPermutation("orly", "roly"))
	fmt.Println(one.CheckPermutation("ooo1123", "1o2o3o1"))
}
