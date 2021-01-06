from collections import defaultdict
from trees_1 import TreeNode

# for TreeNode class defintion see Minimal BST

# Approach:
# write a method that determines if two trees are equivalent by value
# let m be the number of nodes in t1
# let n be the number of nodes in t2
# if we were to determine equivalency for all nodes in t1 whose value
# matched the root of t2, the run time would be O(mn)
# instead fully traverse t1 and fully traverse t2 in O(m + n) and use
# what we learn to be much more efficient about which subtrees we comapare
# specifically, the height of t2 must equal the height of any subtree in t1
# worth running the full equivalency check on


# O(n) where n is the number of nodes in the tree
def traverse_tree(node, depth=0, height_to_nodes=None, target=None):
    # a single node has height = 0, so a null node must have height = -1
    if not node:
        return depth - 1

    # get the height of the left subtree
    # store target nodes and heights along the way
    left_height = traverse_tree(
        node.left, depth + 1, height_to_nodes, target)

    # get the height of the right subtee
    # store target nodes and heights along the way
    right_height = traverse_tree(
        node.right, depth + 1, height_to_nodes, target)

    # height is the height of the original tree for a path through this node
    # or the depth of this node if it is a leaf node
    # this value less the current depth is the height of the subtree
    # and the dictionary key used to store nodes matching the target value
    height = max(left_height, right_height)

    if node.data == target:
        height_to_nodes[height - depth].append(node)

    return height


# O(n) where n is the number of nodes in t2
def value_equivalent(t1, t2):
    if t1 == None and t2 == None:
        return True
    elif t1.data != t2.data:
        return False

    return value_equivalent(t1.left, t2.left) and value_equivalent(t1.right, t2.right)


# O(m + n) where m is the number of nodes in t1 and n is the number of nodes in t2
def check_subtree(t1, t2):
    # O(n)
    height = traverse_tree(t2)
    height_to_nodes = defaultdict(list)
    # O(m)
    traverse_tree(t1, height_to_nodes=height_to_nodes, target=t2.data)

    # assuming the heights of t1 and t2 are log m and log n respectively
    # then we're checking at most k nodes at the log m - log n depth in t1
    # there are at most 2^d nodes at depth d, if d = log m - log n = log(m/n)
    # k = O(2^d) = O(2^log(m/n)) = O(m/n)
    # so this loop runs O(m/n) times where value_equivalent costs O(n)
    # as a note, value_equivalent will likely make far fewer than n
    # recursive calls for the cases where the trees are not equivalent
    # O(m/n) outer loop iterations * O(n) cost per puts this at O(m)
    # assuming reasonably balanced trees
    for node in height_to_nodes[height]:
        # O(n)
        if value_equivalent(node, t2):
            return True

    return False
