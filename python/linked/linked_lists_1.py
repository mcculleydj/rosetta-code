class LinkedList:
    def __init__(self):
        self.head = None

    def __repr__(self):
        node = self.head
        nodes = []
        while node is not None:
            nodes.append(str(node.value))
            node = node.next
        nodes.append('None')
        return " -> ".join(nodes)


class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

    def __repr__(self):
        return str(self.value)


# O(n)
def remove_duplicates_time(linked_list):
    values = set()
    node = linked_list.head

    while node.next != None:
        if node.next.value in values:
            node.next = node.next.next
        else:
            values.add(node.next.value)
            node = node.next


# O(1) space -- O(n^2) time
def remove_duplicates_space(linked_list):
    index_node = linked_list.head

    while index_node.next != None:
        cursor_node = index_node
        while cursor_node.next != None:
            if cursor_node.next.value == index_node.value:
                cursor_node.next = cursor_node.next.next
            else:
                cursor_node = cursor_node.next
        index_node = index_node.next


ll = LinkedList()

n1 = Node(1)
n2 = Node(2)
n3 = Node(3)
n4 = Node(4)
n5 = Node(4)
n6 = Node(5)
n7 = Node(5)

ll.head = n1
n1.next = n2
n2.next = n3
n3.next = n4
n4.next = n5
n5.next = n6
n6.next = n7

print(ll)
remove_duplicates_time(ll)
# remove_duplicates_space(ll)
print(ll)
