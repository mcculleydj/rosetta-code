# see Remove Duplicates for List and Node type definitions


# O(n)
def partition(list, value):
    head = list.head
    tail = list.head
    node = list.head

    # O(n)
    while node:
        # keep a reference to the next node in the original list
        next = node.next

        if node.value < value:
            # reassign node to point at the previous head
            node.next = head
            # reassign the head of the list to this node
            head = node
        else:
            # push this node on to the end of the list
            tail.next = node
            # reassign the tail to this node
            tail = node

        # move on to the next node in the original list
        node = next

    # clip the tail
    tail.next = None

    return head
