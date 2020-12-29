from linked_lists_1 import LinkedList, Node

# see Remove Duplicates for List and Node type definitions


# O(n)
def find_loop_start(l):
    visited = set()
    node = l.head
    while node:
        if node in visited:
            return node
        visited.add(node)
        node = node.next


# let k be the distance between head and the start of the loop
# let s be a slow pointer that moves one node every step
# let f be a fast pointer that moves two nodes every step
# let l be the length of the loop

# when s has moved k and enters the loop
# f is k steps into the loop, since it moved a distance of 2k
# since the loop is a circular path and we know nothing about the length of the loop
# compared to the distance between the head and the start of the loop
# f is actually a distance k % l past the start of the loop when s gets there
# having gone around the loop 0 or more times

# let K = k % l
# once both pointers are in the loop f gets 1 node closer to s every step
# f is l - K steps behind s and the two will collide after exactly l - K steps
# since s was at the start of the loop we can equivalently say that they will
# collide at a distance of K from the start of the loop

# since K = k % l, we can represent this as K + m * l = k for any integer m
# since 1 degree on a circle is the same position as 361 degrees on a circle
# we can say that a distance of K from the start is the same as a distance of k

# knowing that the collision point is a k nodes back from the start of the loop
# and k nodes from the head of the list we can traverse two pointers
# one from head and the other from the node at which f and s collided
# the node at which these two pointers collide must be the start of the loop

def find_loop_start_hard(l):
    fast = l.head
    slow = l.head

    # take the first step (i.e. do while)
    fast = fast.next.next
    slow = slow.next

    # "is" provides memory reference equivalence
    while fast is not slow:
        fast = fast.next.next
        slow = slow.next

    head_cursor = l.head
    collision_cursor = fast

    while head_cursor is not collision_cursor:
        head_cursor = head_cursor.next
        collision_cursor = collision_cursor.next

    return head_cursor


n5 = Node(5, None)
n4 = Node(4, n5)
n3 = Node(3, n4)
n2 = Node(2, n3)
n1 = Node(1, n2)
l = LinkedList(n1)

n5.next = n2

print(find_loop_start_hard(l))
