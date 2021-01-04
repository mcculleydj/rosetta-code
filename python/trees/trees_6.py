from trees_1 import TreeNode

# for TreeNode class defintion see Minimal BST


# O(n)
# recursive scan of every node to determine existence of targets
# given a node and a tuple of target nodes, return a tuple of boolean values
# signifying whether or not each target appears in the tree
def check_for_nodes(node, targets):
    if node == None:
        return False, False
    elif node is targets[0]:
        hasOther = check_for_nodes(node.left, (targets[1],)) or check_for_nodes(
            node.right, (targets[1],))
        return True, hasOther[0]
    elif len(targets) > 1 and node is targets[1]:
        hasOther = check_for_nodes(node.left, (targets[0],)) or check_for_nodes(
            node.right, (targets[0],))
        return hasOther[0], True

    check_left = check_for_nodes(node.left, targets)
    check_right = check_for_nodes(node.right, targets)

    return (check_left[0] or check_right[0]), (check_left[1] or check_right[1])


# O(n log n) for a balanced tree
# O(n^2) worst case
# ---
# O(log n) recursive calls to first_common_ancestor for a balanced tree
# O(n) recursive calls to first_common_ancestor worst case
# O(n) for each call to check for nodes
def first_common_ancestor(node, targets):
    one_in_left, two_in_left = check_for_nodes(node.left, targets)
    one_in_right, two_in_right = check_for_nodes(node.right, targets)

    # if one target is found in one subtree and the other in the other
    # then this is the first common ancestor
    if (one_in_left and two_in_right) or (two_in_left and one_in_right):
        return node
    elif one_in_left and two_in_left:
        return first_common_ancestor(node.left, targets)

    # one_in_right and two_in_right
    return first_common_ancestor(node.right, targets)
