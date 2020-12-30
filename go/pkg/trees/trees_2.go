package trees

import (
	"fmt"
	"strings"
)

// for TreeNode type def see Minimal BST

// QueueNode stores a tree node, its depth, and a pointer to the next node
type QueueNode struct {
	TreeNode *TreeNode
	Depth    int
	Next     *QueueNode
}

// ToString QueueNode => string
func (n *QueueNode) ToString() string {
	return fmt.Sprintf("(%d, %d)", n.TreeNode.Data, n.Depth)
}

// Queue is a FIFO data structure
type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Length int
}

// ToString Queue => string
func (q *Queue) ToString() string {
	ss := []string{}
	node := q.Head

	for node != nil {
		ss = append(ss, node.ToString())
		node = node.Next
	}

	return strings.Join(ss, " -> ")
}

// Add a new node in O(1)
func (q *Queue) Add(treeNode *TreeNode, depth int) {
	queueNode := new(QueueNode)
	queueNode.TreeNode = treeNode
	queueNode.Depth = depth

	if q.Head == nil {
		q.Head = queueNode
	}

	if q.Tail != nil {
		q.Tail.Next = queueNode
	}

	q.Tail = queueNode
	q.Length++
}

// Remove a node and return the tree node and depth in O(1)
func (q *Queue) Remove() (*TreeNode, int) {
	if q.Head == nil {
		panic("queue is empty")
	}

	node := q.Head
	q.Head = q.Head.Next
	q.Length--

	return node.TreeNode, node.Depth
}

// ListDepths builds a slice of queues in O(n)
func ListDepths(treeNode *TreeNode, depth int, q *Queue, depths []*Queue) []*Queue {
	// add an additional queue for each depth encountered
	if depth == len(depths) {
		depths = append(depths, new(Queue))
	}

	// add node to queue on depths slice
	depths[depth].Add(treeNode, depth)

	if treeNode.Left != nil {
		q.Add(treeNode.Left, depth+1)
	}

	if treeNode.Right != nil {
		q.Add(treeNode.Right, depth+1)
	}

	if q.Length == 0 {
		return depths
	}

	treeNode, depth = q.Remove()

	return ListDepths(treeNode, depth, q, depths)
}
