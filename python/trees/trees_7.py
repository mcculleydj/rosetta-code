from trees_1 import TreeNode
from collections import defaultdict

# for TreeNode class defintion see Minimal BST

# instantiate a dictionary int => list
# traverse the tree while keeping track of the current depth
# store all values for a given depth (starting at depth 0)
# construct a series of lists where permutations
# of the lists in the dictionary are concatenated in order
# e.g. 1 => (2 => (4, 5), 3) would result in
# depth 0 permutations: {[1]}
# depth 1 permutations: {[2, 3], [3, 2]}
# depth 2 permutations: {[4, 5], [5, 4]}
# leading to the following set of concatenated lists:
# {
#   [1, 2, 3, 4, 5],
#   [1, 2, 3, 5, 4],
#   [1, 3, 2, 4, 5],
#   [1, 3, 2, 5, 4]
# }


def construct_depth_map(node, depth=0, depth_map=defaultdict(list)):
    if not node:
        return
    depth_map[depth].append(node.data)
    construct_depth_map(node.left, depth + 1, depth_map)
    construct_depth_map(node.right, depth + 1, depth_map)

    return depth_map


def permutations(a, k, acc):
    if k == len(a):
        permutation = a.copy()
        acc.append(permutation)
    else:
        for i in range(k, len(a)):
            # swap the ith and kth elements
            a[k], a[i] = a[i], a[k]
            # call permutations on the rest of the list
            permutations(a, k + 1, acc)
            # swap the elements back to their original position
            # before incrementing i
            a[k], a[i] = a[i], a[k]

    return acc


def bst_sequences(root):
    depth_map = construct_depth_map(root)
    sequences = [[root.data]]

    for d in range(1, len(depth_map)):
        appended_sequences = []
        for p in permutations(depth_map[d], 0, []):
            for s in sequences:
                appended_sequences.append(s + p)
        sequences = appended_sequences

    return sequences
