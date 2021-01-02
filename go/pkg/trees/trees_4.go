package trees

// for TreeNode type def see Minimal BST

// IsBST runs in O(n)
// using pointers for the bounds to leverage
// a nil zero value rather than 0 which
// is valid data for a TreeNode<int>
func IsBST(tn *TreeNode, leftBound, rightBound *int) bool {
	if tn == nil {
		return true
	}

	if leftBound != nil && tn.Data <= *leftBound {
		return false
	}

	if rightBound != nil && tn.Data <= *rightBound {
		return false
	}

	return IsBST(tn.Left, leftBound, &tn.Data) && IsBST(tn.Right, &tn.Data, rightBound)
}
