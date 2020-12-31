package trees

import (
	"math"
	"sync"
)

// for TreeNode type def see Minimal BST

// this implementation builds on the Python solution
// to exhibit some of what makes Go a wonderful tool
// for dealing with concurrency

// NodeState represents the state of a node in the tree
// with respect to the properties we care about:
// height and balance of the two subtrees
// height can only be an int, but Go's math module requires float64
// for Max() and Abs()
type NodeState struct {
	IsBalanced bool
	Height     float64
}

// IsBalanced runs in O(n)
func IsBalanced(tn *TreeNode, wgs chan *sync.WaitGroup, wg *sync.WaitGroup, result *NodeState) {
	// defer an anonymous function so that
	// the WaitGroup is always completed regardless of how this function exits
	// without defer this code would need to appear above every return statement
	// hooray Go!
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	// base case: nil node
	// by convention a single node has height 0, hence -1 for a nil
	// set NodeState and return
	if tn == nil {
		result.IsBalanced = true
		result.Height = -1
		return
	}

	// allocate two new NodeStates
	left := new(NodeState)
	right := new(NodeState)

	// select is one my favorite Go concepts
	// if a WaitGroup is available this thread is free to divide
	// otherwise proceed synchronously
	select {
	case wgNext := <-wgs:
		// set up a stop condition until wgNext.Done() is called 2 times
		wgNext.Add(2)

		// launch two new threads in parallel to check the right and left subtrees
		go IsBalanced(tn.Left, wgs, wgNext, left)
		go IsBalanced(tn.Right, wgs, wgNext, right)

		// wait for Done() to be called 2 times
		wgNext.Wait()

		// put this WaitGroup back on the channel for another thread to use
		wgs <- wgNext
	default:
		// no bandwidth left to split, proceed without concurrency
		IsBalanced(tn.Left, wgs, nil, left)

		// since this is sequential we can save a call on the right subtree
		// if the left comes back false
		if !left.IsBalanced {
			result.IsBalanced = false
			return
		}

		IsBalanced(tn.Right, wgs, nil, right)
	}

	// check the imbalance condition at this node
	// now that we know the heights of the subtrees
	if math.Abs(left.Height-right.Height) > 1 {
		result.IsBalanced = false
		return
	}

	// carry the IsBalanced state of all decendants up to the root
	result.IsBalanced = left.IsBalanced && right.IsBalanced

	// height is the maximum height of the left and right subtree plus this node
	result.Height = 1 + math.Max(left.Height, right.Height)
}

// IsBalancedWrapper simply returns the boolean portion of the struct in O(n)
// this function also sets up the concurrency infrastructure
func IsBalancedWrapper(tn *TreeNode) bool {
	// there can exist up to maxSplits * 2 threads
	// since a single waitGroup is being used to support two parallel recursive calls
	const maxSplits = 4
	wgs := make(chan *sync.WaitGroup, maxSplits)

	for i := 0; i < maxSplits; i++ {
		wgs <- new(sync.WaitGroup)
	}

	result := new(NodeState)
	IsBalanced(tn, wgs, nil, result)

	return result.IsBalanced
}
