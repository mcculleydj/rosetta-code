package linked

// see Remove Duplicates for List and Node type definitions

// DeleteNode is O(1) -- see Python for comments
func DeleteNode(n *Node) {
	n.Value = n.Next.Value
	n.Next = n.Next.Next
}
