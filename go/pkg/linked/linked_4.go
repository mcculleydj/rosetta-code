package linked

// see Remove Duplicates for List and Node type definitions

// Partition runs in O(n)
func Partition(l List, value int) *Node {
	head := l.Head
	tail := l.Head
	node := l.Head

	for node != nil {
		next := node.Next

		if node.Value < value {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}

		node = next
	}

	tail.Next = nil

	return head
}
