package trees

import (
	"fmt"
	"strconv"
)

// TreeNodeWithParent represents a binary tree node with a reference to its parent node
type TreeNodeWithParent struct {
	Data   int
	Parent *TreeNodeWithParent
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
}

// ToString provides a string representation for binary tree node
func (n *TreeNodeWithParent) ToString() string {
	if n == nil {
		return "nil"
	}
	if n.Left == nil && n.Right == nil {
		return strconv.Itoa(n.Data)
	}
	return fmt.Sprintf("%d => (%s, %s)", n.Data, n.Left.ToString(), n.Right.ToString())
}

// checkRight runs in O(log n)
func checkRight(tn *TreeNodeWithParent) *TreeNodeWithParent {
	for tn.Left != nil {
		tn = tn.Left
	}

	return tn
}

// checkParent runs in O(log n)
func checkParent(tn *TreeNodeWithParent) *TreeNodeWithParent {
	parent := tn.Parent

	for parent != nil {
		if parent.Data >= tn.Data {
			return parent
		}
		parent = parent.Parent
	}

	return nil
}

// FindNext runs in O(log n)
func FindNext(tn *TreeNodeWithParent) *TreeNodeWithParent {
	if tn == nil {
		return nil
	}

	if tn.Right != nil {
		return checkRight(tn.Right)
	}

	return checkParent(tn)
}