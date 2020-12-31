package trees

import (
	"fmt"
	"strconv"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// ToString provides a string representation for binary tree node
func (n *TreeNode) ToString() string {
	if n == nil {
		return "nil"
	}
	if n.Left == nil && n.Right == nil {
		return strconv.Itoa(n.Data)
	}
	return fmt.Sprintf("%d => (%s, %s)", n.Data, n.Left.ToString(), n.Right.ToString())
}

// O(1)
func splitSlice(ns []int) (middle int, left []int, right []int) {
	i := len(ns) / 2
	return ns[i], ns[:i], ns[i+1:]
}

// MinBST runs in O(n)
func MinBST(ns []int, node *TreeNode) *TreeNode {
	if len(ns) == 0 {
		return nil
	}

	if node == nil {
		node = new(TreeNode)
	}

	data, left, right := splitSlice(ns)
	node.Data = data

	if len(left) > 0 {
		node.Left = new(TreeNode)
		MinBST(left, node.Left)
	}

	if len(right) > 0 {
		node.Right = new(TreeNode)
		MinBST(right, node.Right)
	}

	return node
}
