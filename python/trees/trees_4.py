from trees_1 import TreeNode

# for tree node class definition see Minimal BST

# base case
# None => True

# recursive case
# node.data must fall within [left bound, right bound)
# left and right bounds are defined by the min / max ancestors
# left and right subtrees must satisfy isBST()


# O(n)
def isBST(tn, left_bound=None, right_bound=None):
    if not tn:
        return True

    # must be strictly greater than any ancestor to the left
    if left_bound is not None and tn.data <= left_bound:
        return False

    # must be less than or equal to any ancestor to the right
    if right_bound is not None and tn.data > right_bound:
        return False

    return isBST(tn.left, left_bound, tn.data) and isBST(tn.right, tn.data, right_bound)
