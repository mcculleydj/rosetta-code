class TreeNode:
    def __init__(self, data=None, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

    def __repr__(self):
        if not self.left and not self.right:
            return str(self.data)
        return f'{self.data} => ({self.left}, {self.right})'


# O(1) return the middle of the list and a reference
# to left and right sublists not including the middle item
def splice_middle(l):
    i = len(l) // 2
    return l[i], l[:i], l[i+1:]


# O(n) one recursive call per item in the list
def min_bst(l, node=TreeNode()):
    if len(l) == 0:
        return

    data, left, right = splice_middle(l)
    node.data = data

    if len(left):
        node.left = TreeNode()
        min_bst(left, node.left)
    if len(right):
        node.right = TreeNode()
        min_bst(right, node.right)

    return node
