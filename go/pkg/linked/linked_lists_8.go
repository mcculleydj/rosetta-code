package linked

// FindLoopStart is O(n)
func FindLoopStart(l List) *Node {
	m := map[*Node]bool{}
	node := l.Head

	for node != nil {
		if m[node] {
			return node
		}
		m[node] = true
		node = node.Next
	}

	return nil
}

// FindLoopStartHard is O(n)
// see comments in Python solution
func FindLoopStartHard(l List) *Node {
	fast := l.Head
	slow := l.Head

	fast = fast.Next.Next
	slow = slow.Next

	for fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = l.Head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
