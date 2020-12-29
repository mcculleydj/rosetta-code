package linked

// FindIntersection is O(m + n)
func FindIntersection(l1, l2 List) *Node {
	m := map[*Node]bool{}

	node := l1.Head
	for node != nil {
		m[node] = true
		node = node.Next
	}

	node = l2.Head
	for node != nil {
		if m[node] {
			return node
		}
		node = node.Next
	}

	return nil
}
