package trees

import "math"

// see Minimal BST for TreeNode type def

// O(n)
func traverseTree(tn *TreeNode, depth int, heightToNodes map[int][]*TreeNode, target *int) int {
	if tn == nil {
		return depth - 1
	}

	leftHeight := traverseTree(tn.Left, depth+1, heightToNodes, target)
	rightHeight := traverseTree(tn.Right, depth+1, heightToNodes, target)

	height := int(math.Max(float64(leftHeight), float64(rightHeight)))

	if target != nil && tn.Data == *target {
		heightToNodes[height-depth] = append(heightToNodes[height-depth], tn)
	}

	return height
}

// O(n)
func valueEquivalent(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1.Data != t2.Data {
		return false
	}

	return valueEquivalent(t1.Left, t2.Left) && valueEquivalent(t1.Right, t2.Right)
}

// CheckSubtree runs in O(m + n)
func CheckSubtree(t1, t2 *TreeNode) bool {
	height := traverseTree(t2, 0, nil, nil)
	heightToNodes := map[int][]*TreeNode{}
	traverseTree(t1, 0, heightToNodes, &t2.Data)

	for _, node := range heightToNodes[height] {
		if valueEquivalent(node, t2) {
			return true
		}
	}

	return false
}
