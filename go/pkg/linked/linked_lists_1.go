package linked

import "fmt"

// Node is a node in a linked list
// value could be any datatype, but Go does not support generics
type Node struct {
	Value int
	Next  *Node
}

// List wraps a head node
type List struct {
	Head *Node
}

// ToString prints a linked list
func (l List) ToString() string {
	s := ""
	node := l.Head
	for node != nil {
		s = s + fmt.Sprintf("%d -> ", node.Value)
		node = node.Next
	}
	return s + "nil"
}

// RemoveDuplicatesTime runs in O(n)
func RemoveDuplicatesTime(l List) {
	valueSet := map[int]bool{}
	node := l.Head

	for node.Next != nil {
		if valueSet[node.Next.Value] {
			node.Next = node.Next.Next
		} else {
			valueSet[node.Next.Value] = true
			node = node.Next
		}
	}
}

// RemoveDuplicatesSpace runs in O(1) space -- O(n^2) time
func RemoveDuplicatesSpace(l List) {
	indexNode := l.Head
	var cursorNode *Node

	for indexNode.Next != nil {
		cursorNode = indexNode
		for cursorNode.Next != nil {
			if cursorNode.Next.Value == indexNode.Value {
				cursorNode.Next = cursorNode.Next.Next
			} else {
				cursorNode = cursorNode.Next
			}
		}
		indexNode = indexNode.Next
	}
}
