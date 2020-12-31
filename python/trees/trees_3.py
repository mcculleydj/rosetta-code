from trees_1 import TreeNode

# for tree node class definition see Minimal BST


# O(n)
# figures out the height of the tree
# while looking for imbalance at every node
def is_balanced(node):
    if not node:
        return (True, -1)

    left_height = is_balanced(node.left)

    if not left_height[0]:
        return left_height

    right_height = is_balanced(node.right)

    if not right_height[0]:
        return right_height

    if abs(left_height[1] - right_height[1]) > 1:
        # the number part of the tuple no longer matters
        return (False, 0)

    return (True, 1 + max(left_height[1], right_height[1]))


# O(n)
def is_balanced_wrapper(node):
    return is_balanced(node)[0]
