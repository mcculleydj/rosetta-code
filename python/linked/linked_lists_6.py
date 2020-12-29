from linked_lists_1 import LinkedList, Node

# see Remove Duplicates for List and Node type definitions


# O(n)
def list_to_stack(l):
    stack = []
    node = l.head

    while node:
        stack.append(node)
        node = node.next

    return stack


# O(n)
def is_palindrome(l):
    list_node = l.head
    # O(n) create a stack to represent the list in reverse order
    stack = list_to_stack(l)
    list_length = len(stack)

    # O(n) compare the stack to the list
    while len(stack) >= list_length // 2:
        stack_node = stack.pop()
        if stack_node.value != list_node.value:
            return False
        list_node = list_node.next

    return True
