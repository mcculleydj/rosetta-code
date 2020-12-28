# see Remove Duplicates for List and Node type definitions


# O(1)
def delete_node(node):
    # overwrite the value of this node with the next node
    # then splice out the next node to remove the duplicate data

    # assert that node.next exists from the problem statement
    node.value = node.next.value
    node.next = node.next.next

    # note that this does not remove the reference to the node
    # in fact it removes the reference to the next node
    # it only ensures that the data in the list reflects the delete
