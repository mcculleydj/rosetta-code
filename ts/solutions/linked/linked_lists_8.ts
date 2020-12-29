import { LinkedList, Node_ } from './linked_lists_1'

// findLoopStart is O(n)
function findLoopStart(l: LinkedList): Node_ {
  const visited: Set<Node_> = new Set()
  let node = l.head

  while (node) {
    if (visited.has(node)) {
      return node
    }
    visited.add(node)
    node = node.next
  }
}

// findLoopStartHard is O(n)
// see comments in Python solution
function findLoopStartHard(l: LinkedList): Node_ {
  let fast = l.head
  let slow = l.head

  fast = fast.next.next
  slow = slow.next

  while (fast !== slow) {
    fast = fast.next.next
    slow = slow.next
  }

  slow = l.head

  while (fast !== slow) {
    fast = fast.next
    slow = slow.next
  }

  return fast
}
