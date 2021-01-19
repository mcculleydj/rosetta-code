package trees

// for TreeNode type def see Minimal BST

// O(1)
func updatePaths(paths map[int]int, key, delta int) {
	// this works because paths is a ptr to the underlying map struct
	// however, Go does not have reference variables in the C++ sense
	// if I reassign paths to a new map in this closure
	// the caller will still maintain its reference to the original map
	// Go omits the * in the signature because it was always
	// being written by early Go developers

	updatedValue := paths[key] + delta

	if updatedValue == 0 {
		delete(paths, key)
	} else {
		paths[key] = updatedValue
	}
}

// PathsToSum runs in O(n)
func PathsToSum(node *TreeNode, target, current int, paths map[int]int) int {
	if node == nil {
		return 0
	}

	current = current + node.Data
	total := paths[current-target]

	if current == target {
		total++
	}

	updatePaths(paths, current, 1)
	total = total + PathsToSum(node.Left, target, current, paths)
	total = total + PathsToSum(node.Right, target, current, paths)
	updatePaths(paths, current, -1)

	return total
}
