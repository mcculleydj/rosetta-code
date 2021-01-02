class TreeNodeWithParent:
    def __init__(self, data=None, parent=None, left=None, right=None):
        self.data = data
        self.parent = parent
        self.left = left
        self.right = right

    def __repr__(self):
        if not self.left and not self.right:
            return str(self.data)
        return f'{self.data} => ({self.left}, {self.right})'


# O(log n)
# return the left most node of this subtree
def check_right(tn):
    while tn.left:
        tn = tn.left

    return tn


# O(log n)
# return the first ancestor whose value is greater than
# or equal to the value of the source node
def check_parent(tn):
    parent = tn.parent

    while parent:
        if parent.data >= tn.data:
            return parent
        parent = parent.parent

    return None


def find_next(tn):
    if not tn:
        return None

    if tn.right:
        return check_right(tn.right)

    return check_parent(tn)
