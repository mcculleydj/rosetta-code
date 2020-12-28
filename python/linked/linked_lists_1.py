class LinkedList:
    def __init__(self, head):
        self.head = head

    def __repr__(self):
        node = self.head
        nodes = []
        while node:
            nodes.append(str(node.value))
            node = node.next
        return ' -> '.join(nodes + ['None'])


class Node:
    def __init__(self, value, next):
        self.value = value
        self.next = next

    def __repr__(self):
        return str(self.value)


# O(n)
def remove_duplicates_time(linked_list):
    values = set()
    node = linked_list.head

    while node.next:
        if node.next.value in values:
            node.next = node.next.next
        else:
            values.add(node.next.value)
            node = node.next


# O(1) space -- O(n^2) time
def remove_duplicates_space(linked_list):
    index_node = linked_list.head

    while index_node.next:
        cursor_node = index_node
        while cursor_node.next:
            if cursor_node.next.value == index_node.value:
                cursor_node.next = cursor_node.next.next
            else:
                cursor_node = cursor_node.next
        index_node = index_node.next
