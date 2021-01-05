package trees

// see Minimal BST for TreeNode type def

func constructDepthMap(tn *TreeNode, depth int, depthMap map[int][]int) {
	if tn == nil {
		return
	}

	depthMap[depth] = append(depthMap[depth], tn.Data)
	constructDepthMap(tn.Left, depth+1, depthMap)
	constructDepthMap(tn.Right, depth+1, depthMap)
}

func copy(ns []int) []int {
	clone := []int{}

	for _, n := range ns {
		clone = append(clone, n)
	}

	return clone
}

func swap(ns []int, i, j int) {
	temp := ns[i]
	ns[i] = ns[j]
	ns[j] = temp
}

func permutations(ns []int, k int, acc *[][]int) {
	if k == len(ns) {
		*acc = append(*acc, copy(ns))
	} else {
		for i := k; i < len(ns); i++ {
			swap(ns, i, k)
			permutations(ns, k+1, acc)
			swap(ns, i, k)
		}
	}
}

// BSTSequences permutes each layer of the tree
// to generate the sequences that could've produced it
func BSTSequences(root *TreeNode) [][]int {
	depthMap := map[int][]int{}
	constructDepthMap(root, 0, depthMap)
	sequences := [][]int{[]int{root.Data}}

	for d := 1; d < len(depthMap); d++ {
		appendedSequences := [][]int{}
		acc := [][]int{}
		permutations(depthMap[d], 0, &acc)

		for _, p := range acc {
			for _, s := range sequences {
				sequence := append(s, p...)
				appendedSequences = append(appendedSequences, sequence)
			}
		}

		sequences = appendedSequences
	}

	return sequences
}
