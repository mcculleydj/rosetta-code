from trees_1 import TreeNode

# for tree node class definition see Minimal BST

# Approach
# traverse using DFS
# for each node track a current sum
# as well as the current sums of each ancestor back to the root
# mapped to the number of times that current sum appears
#     0
#   0
# 3  3
# would yield a dictionary with {0: 2} at each leaf node
# this dictionary is updated before and after the recursion
# current sums are incremented before and decremented after
# so that the right subtree does not use a dictionary
# based on values in the left subtree
# the only other case to handle is when the current sum is the target
# where the root of this tree alone would be a valid path based
# on the current sum of the parent node (as in the case of both leaf nodes above)
# in this case, increment the total once before adding it to the recursive calls


# O(1)
def update_paths(paths, key, delta):
    new = paths.get(key, 0) + delta

    if new == 0:
        del paths[key]
    else:
        paths[key] = new


# O(n)
def paths_to_sum(node, target, current, paths):
    if node is None:
        return 0

    current += node.data
    total = paths.get(current - target, 0)

    # root of this subtree is a valid path based on current sum
    if current == target:
        total += 1

    update_paths(paths, current, 1)
    total += paths_to_sum(node.left, target, current, paths)
    total += paths_to_sum(node.right, target, current, paths)
    update_paths(paths, current, -1)

    return total
