from trees_1 import Node as TreeNode

# for tree node class definition see Minimal BST


class QueueNode:
    def __init__(self, tree_node=None, depth=None, next=None):
        self.tree_node = tree_node
        self.depth = depth
        self.next = next

    def __repr__(self):
        return str((self.tree_node.data, self.depth))


# implemented for O(1) removal of the first element
class Queue:
    def __init__(self):
        self.tail = None
        self.head = None
        self.length = 0

    def __repr__(self):
        nodes = []
        node = self.head

        while node:
            nodes.append(str(node))
            node = node.next

        return ' -> '.join(nodes)

    def add(self, tree_node, depth):
        node = QueueNode(tree_node, depth)

        if not self.head:
            self.head = node

        if self.tail:
            self.tail.next = node

        self.tail = node
        self.length += 1

    def remove(self):
        if not self.head:
            raise Exception('queue is empty')

        node = self.head
        self.head = self.head.next
        self.length -= 1

        return node.tree_node, node.depth


# O(n)
# node_depth is a tuple containing the binary tree node and the depth it was found
# queue is used to implement BFS storing (node, depth) tuples for FIFO access
# depths is a list of queues containing all node data at depth d
# d is both the index of the depths list and the depth in the tree
def list_depths(tree_node, depth, queue=Queue(), depths=[]):
    if depth == len(depths):
        depths.append(Queue())

    depths[depth].add(tree_node, depth)

    if tree_node.left:
        queue.add(tree_node.left, depth + 1)
    if tree_node.right:
        queue.add(tree_node.right, depth + 1)

    if queue.length == 0:
        return depths

    tree_node, depth = queue.remove()
    return list_depths(tree_node, depth, queue, depths)
