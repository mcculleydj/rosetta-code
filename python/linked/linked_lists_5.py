from linked_lists_1 import LinkedList, Node
# see Remove Duplicates for List and Node type definitions


# O(m + n)
def right_to_left_sum(l1, l2):
    # step through the lists at the same time performing addition
    # and carrying the one as needed
    n1 = l1.head
    n2 = l2.head
    # assumes length of each list is at least 1
    next_value = n1.value + n2.value
    carry = next_value // 10
    sum_node = Node(next_value % 10, None)
    sum_list = LinkedList(sum_node)

    # O(m + n)
    while n1.next or n2.next:
        if not n1.next:
            next_value = n2.next.value + carry
        elif not n2.next:
            next_value = n1.next.value + carry
        else:
            next_value = n1.next.value + n2.next.value + carry

        carry = next_value // 10

        sum_node.next = Node(next_value % 10, None)
        sum_node = sum_node.next

        n1 = n1.next if n1.next else n1
        n2 = n2.next if n2.next else n2

    if carry == 1:
        sum_node.next = Node(1, None)

    return sum_list


# O(n)
def list_to_stack(l):
    stack = []
    node = l.head
    while node:
        stack.append(node)
        node = node.next

    return stack


# O(m + n)
def left_to_right_sum(l1, l2):
    # O(n)
    stack1 = list_to_stack(l1)
    # O(m)
    stack2 = list_to_stack(l2)

    head = None
    carry = 0

    # O(m + n)
    while stack1 or stack2:
        if not stack1:
            node = stack2.pop()
            next_value = node.value + carry
            node.value = next_value % 10
            node.next = head
        elif not stack2:
            node = stack1.pop()
            next_value = node.value + carry
            node.value = next_value % 10
            node.next = head
        else:
            node1 = stack1.pop()
            node2 = stack2.pop()
            next_value = node1.value + node2.value + carry
            node = Node(next_value % 10, head)

        carry = next_value // 10
        head = node

    if carry == 1:
        node = Node(1, head)
        head = node

    return LinkedList(head)


# n4 = Node(1, None)
# n3 = Node(2, n4)
# n2 = Node(9, n3)
# n1 = Node(5, n2)
# l1 = LinkedList(n1)

# n7 = Node(8, None)
# n6 = Node(5, n7)
# n5 = Node(3, n6)
# l2 = LinkedList(n5)

n4 = Node(9, None)
n3 = Node(9, n4)
n2 = Node(9, n3)
n1 = Node(9, n2)
l1 = LinkedList(n1)

n5 = Node(1, None)
l2 = LinkedList(n5)


# print(right_to_left_sum(l1, l2))
print(left_to_right_sum(l1, l2))
