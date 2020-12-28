package linked

import "fmt"

// see Remove Duplicates for List and Node type definitions

// KthToLastElement runs in O(n)
func KthToLastElement(l List, k int) *Node {
	leadNode := l.Head
	lagNode := l.Head
	leadIndex := 0
	lagIndex := 0

	for leadNode != nil {
		leadNode = leadNode.Next
		leadIndex++
		if leadIndex > k {
			lagNode = lagNode.Next
			lagIndex++
		}
	}

	if leadIndex-lagIndex < k {
		fmt.Printf("List is not long enough to have a %d(st/nd/rd/th) to last element, returning the first node instead.\n", k)
	}

	return lagNode
}
