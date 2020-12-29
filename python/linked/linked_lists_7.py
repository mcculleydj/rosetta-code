from linked_lists_1 import LinkedList, Node

# see Remove Duplicates for List and Node type definitions


# O(m + n) time -- requires additional space for the set
def find_intersection(l1, l2):
    # could do this in O(m + n) without the set
    # by traversing to determine length of each and if the tail nodes are the same
    # then traverse both lists starting at the difference between the lengths for the longer list
    # effectively have the same distance between the tail and the starting point for each list
    # compare each next node and the first node that is the same is the intersection

    list1_nodes = set()

    node = l1.head
    while node:
        list1_nodes.add(node)
        node = node.next

    node = l2.head
    while node:
        if node in list1_nodes:
            return node
        node = node.next


n3 = Node(3, None)
n2 = Node(2, n3)
n1 = Node(1, n2)
l1 = LinkedList(n1)


n5 = Node(5, n2)
n4 = Node(4, n5)
l2 = LinkedList(n4)


print(find_intersection(l1, l2))
