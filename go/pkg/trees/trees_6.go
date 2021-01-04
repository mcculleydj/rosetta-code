package trees

// see Minimal BST for TreeNode type def

// O(n)
func checkForNodes(tn, targetOne, targetTwo *TreeNode) (bool, bool) {
	if tn == nil {
		return false, false
	} else if tn == targetOne {
		_, leftHasTwo := checkForNodes(tn.Left, targetOne, targetTwo)
		_, rightHasTwo := checkForNodes(tn.Right, targetOne, targetTwo)
		return true, leftHasTwo || rightHasTwo
	} else if tn == targetTwo {
		leftHasOne, _ := checkForNodes(tn.Left, targetOne, targetTwo)
		rightHasOne, _ := checkForNodes(tn.Right, targetOne, targetTwo)
		return leftHasOne || rightHasOne, true
	}

	leftHasOne, leftHasTwo := checkForNodes(tn.Left, targetOne, targetTwo)
	rightHasOne, rightHasTwo := checkForNodes(tn.Right, targetOne, targetTwo)

	return leftHasOne || rightHasOne, leftHasTwo || rightHasTwo
}

// FirstCommonAncestor runs in O(n^2) -- O(n log n) for a balanced tree
func FirstCommonAncestor(tn, targetOne, targetTwo *TreeNode) *TreeNode {
	leftHasOne, leftHasTwo := checkForNodes(tn.Left, targetOne, targetTwo)
	rightHasOne, rightHasTwo := checkForNodes(tn.Right, targetOne, targetTwo)

	if (leftHasOne && rightHasTwo) || (leftHasTwo && rightHasOne) {
		return tn
	} else if leftHasOne && leftHasTwo {
		return FirstCommonAncestor(tn.Left, targetOne, targetTwo)
	}

	return FirstCommonAncestor(tn.Right, targetOne, targetTwo)
}
