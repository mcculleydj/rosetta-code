package trees

import (
	"fmt"
	"math/rand"
	"time"
)

// BSTNode is a node in a binary search tree
type BSTNode struct {
	Data      int
	Left      *BSTNode
	Right     *BSTNode
	LeftSize  int
	RightSize int
}

// ToString produces a string representation of the tree at n
func (n *BSTNode) ToString() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d | %d | %d => (%s, %s)", n.Data, n.LeftSize, n.RightSize, n.Left.ToString(), n.Right.ToString())
}

// BST is a wrapper around the root node
type BST struct {
	Root *BSTNode
}

// ToString produces a string representation of the tree
func (t *BST) ToString() string {
	return t.Root.ToString()
}

// Insert a value into the tree
func (t *BST) Insert(data int) {
	if t.Root == nil {
		t.Root = &BSTNode{Data: data}
		return
	}

	node := t.Root
	var parent *BSTNode
	var direction string

	for node != nil {
		parent = node

		if data <= node.Data {
			parent.LeftSize++
			node = node.Left
			direction = "left"
		} else {
			parent.RightSize++
			node = node.Right
			direction = "right"
		}
	}

	if direction == "left" {
		parent.Left = &BSTNode{Data: data}
	} else {
		parent.Right = &BSTNode{Data: data}
	}
}

// Find a value in the tree
// return the found node or nil, the parent or nil, and the direction
// from the parent to the found node
func (t *BST) Find(data, decrement int) (*BSTNode, *BSTNode, string) {
	// empty tree
	if t.Root == nil {
		return nil, nil, ""
	}

	// root value is the value we're looking for
	if t.Root.Data == data {
		return t.Root, nil, ""
	}

	node := t.Root
	var parent *BSTNode
	var direction string

	leftDecrementedNodes := []*BSTNode{}
	rightDecrementedNodes := []*BSTNode{}

	for node != nil && node.Data != data {
		parent = node

		if data <= node.Data {
			parent.LeftSize = parent.LeftSize - decrement
			leftDecrementedNodes = append(leftDecrementedNodes, parent)
			node = node.Left
			direction = "left"
		} else {
			parent.RightSize = parent.RightSize - decrement
			rightDecrementedNodes = append(rightDecrementedNodes, parent)
			node = node.Right
			direction = "right"
		}
	}

	if node == nil && decrement > 0 {
		for _, decrementedNode := range leftDecrementedNodes {
			decrementedNode.LeftSize++
		}
		for _, decrementedNode := range rightDecrementedNodes {
			decrementedNode.RightSize++
		}
	}

	return node, parent, direction
}

// Delete a value from the tree
func (t *BST) Delete(data int) {
	node, parent, direction := t.Find(data, 1)

	if node == nil {
		return
	}

	var nodeToShift *BSTNode

	if node.Left != nil && node.Right != nil {
		nodeToShift = node.Left
		var parentOfNodeToShift *BSTNode

		for nodeToShift.Right != nil {
			parentOfNodeToShift = nodeToShift
			parentOfNodeToShift.RightSize--
			nodeToShift = nodeToShift.Right
		}

		// nodeToShift can only have a left subtree after the loop
		parentOfNodeToShift.Right = nodeToShift.Left
		nodeToShift.LeftSize = node.LeftSize - 1
		nodeToShift.RightSize = node.RightSize
	} else if node.Left != nil {
		nodeToShift = node.Left
		nodeToShift.LeftSize = node.LeftSize - 1
	} else if node.Right != nil {
		nodeToShift = node.Right
		nodeToShift.RightSize = node.RightSize - 1
	}
	// unwritten else case: leaf node (Left == nil && Right == nil)
	// safe to simply remove the node in this case

	// fix the size properties and ptrs on the parent
	if parent != nil {
		if direction == "left" {
			parent.Left = nodeToShift
		} else {
			parent.Right = nodeToShift
		}
	} else {
		// handle special case where the root is deleted
		t.Root = nodeToShift
	}

	if nodeToShift != nil {
		// set the children of the shifted node
		// to the deleted node it replaced
		// unless they are equivalent and would cause
		if node.Left != nodeToShift {
			nodeToShift.Left = node.Left
		}
		if node.Right != nodeToShift {
			nodeToShift.Right = node.Right
		}
	}
}

// Random returns a node uniformly at random from the tree
func (t *BST) Random() *BSTNode {
	if t.Root == nil {
		return nil
	}

	node := t.Root
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for node != nil {
		n := r.Float64()
		total := float64(1 + node.LeftSize + node.RightSize)

		if n <= 1/total {
			return node
		} else if n <= float64(1+node.LeftSize)/total {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// should be unreachable
	// could optionally panic here
	return nil
}
