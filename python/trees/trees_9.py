from random import random

# Approach:
# assuming that the data access patterns
# are such that any of the methods are as
# likely to be called as any other we should
# optimize for quick reads and writes
# one way to do this is to create a binary search tree

# for the implementation of random
# as long as each node knows how many nodes
# are in the left and right subtrees
# we can branch the probability of selection
# between left, right and stay accordingly


class BSTNode:
    def __init__(self, data=None):
        self.data = data
        self.left = None
        self.right = None
        self.left_cardinality = 0
        self.right_cardinality = 0

    def __repr__(self):
        return f'{self.data} | {self.left_cardinality} | {self.right_cardinality} => ({self.left}, {self.right})'


class BST:
    def __init__(self, root=None):
        self.root = root

    def __repr__(self):
        return str(self.root)

    # O(log n) generally -- O(n) worst case
    def insert(self, data):
        if not self.root:
            self.root = BSTNode(data)
            return

        node = self.root

        while node:
            parent = node

            if data <= node.data:
                parent.left_cardinality += 1
                node = node.left
                direction = 'left'
            else:
                parent.right_cardinality += 1
                node = node.right
                direction = 'right'

        if direction == 'left':
            parent.left = BSTNode(data)
        else:
            parent.right = BSTNode(data)

    # O(log n) generally -- O(n) worst case
    # finds first node matching data value
    # provides an argument allowing delete to decrement
    def find(self, data, decrement=0):
        node = self.root

        if data == node.data:
            return node, None, None

        left_decremented_nodes = []
        right_decremented_nodes = []

        while node and node.data != data:
            parent = node
            if data <= node.data:
                parent.left_cardinality -= decrement
                left_decremented_nodes.append(parent)
                node = node.left
                direction = 'left'
            else:
                parent.right_cardinality -= decrement
                right_decremented_nodes.append(parent)
                node = node.right
                direction = 'right'

        # if a node to delete is not found
        # need to fix cardinalities
        if not node and decrement > 0:
            for decremented_node in left_decremented_nodes:
                decremented_node.left_cardinality += 1
            for decremented_node in right_decremented_nodes:
                decremented_node.right_cardinality += 1

        # will return None if data value does not exist in tree
        return node, parent, direction

    # O(log n) generally -- O(n) worst case
    def delete(self, data):
        node, parent, direction = self.find(data, decrement=1)
        node_to_shift = None

        # return None if data value is not found
        if not node:
            return

        if node.left and node.right:
            # find right most node of the left subtree
            node_to_shift = node.left

            while node_to_shift.right:
                parent_of_node_to_shift = node_to_shift
                parent_of_node_to_shift.right_cardinality -= 1
                node_to_shift = node_to_shift.right

            # place the left subtree of the shifting node in this node's place
            parent_of_node_to_shift.right = node_to_shift.left
            node_to_shift.left_cardinality = node.left_cardinality - 1
            node_to_shift.right_cardinality = node.right_cardinality
        elif node.left:

            # shift subtree up and to the right
            # in this case the parent is the node being deleted
            node_to_shift = node.left
            node_to_shift.left_cardinality = node.left_cardinality - 1
        elif node.right:
            # shift subtree up and to the left
            # in this case the parent is the node being deleted
            node_to_shift = node.right
            node_to_shift.right_cardinality = node.right_cardinality - 1

        # link parent to shifted subtree depending on direction found
        if parent:
            if direction == 'left':
                parent.left = node_to_shift
            else:
                parent.right = node_to_shift
        else:
            self.root = node_to_shift

        if node_to_shift:
            # link node_to_shift to original node's children
            # provided they are not the same
            if node.left is not node_to_shift:
                node_to_shift.left = node.left
            if node.right is not node_to_shift:
                node_to_shift.right = node.right

    def random(self):
        node = self.root

        # empty tree => None
        if not node:
            return

        while node:
            # a random float between 0.0 and 1.0
            n = random()
            total = 1 + node.left_cardinality + node.right_cardinality

            if n <= 1 / total:
                return node
            elif n <= (1 + node.left_cardinality) / total:
                node = node.left
            else:
                node = node.right
