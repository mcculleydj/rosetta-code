package main

import (
	"fmt"
	"rosetta-code/pkg/graphs"
)

func main() {
	ps := []string{"a", "b", "c"}
	ds := []graphs.Dependency{
		{Blocker: "a", Project: "b"},
		{Blocker: "b", Project: "c"},
	}

	fmt.Println(graphs.BuildSchedule(ps, ds))

	ps = []string{"a", "b", "c", "d", "e", "f"}
	ds = []graphs.Dependency{
		{Blocker: "a", Project: "d"},
		{Blocker: "f", Project: "b"},
		{Blocker: "b", Project: "d"},
		{Blocker: "f", Project: "a"},
		{Blocker: "d", Project: "c"},
	}

	fmt.Println(graphs.BuildSchedule(ps, ds))

	ps = []string{"a", "b", "c"}
	ds = []graphs.Dependency{
		{Blocker: "a", Project: "b"},
		{Blocker: "b", Project: "c"},
		{Blocker: "c", Project: "a"},
	}

	fmt.Println(graphs.BuildSchedule(ps, ds))
}
