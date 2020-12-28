import { LinkedList, Node_ } from './linked_lists_1'

// see Remove Duplicates for List and Node type definitions

// O(n)
function partition(l: LinkedList, value: number): Node_ {
  let head: Node_ = l.head
  let tail: Node_ = l.head
  let node: Node_ = l.head

  while (node) {
    const next = node.next

    if (node.value < value) {
      node.next = head
      head = node
    } else {
      tail.next = node
      tail = node
    }

    node = next
  }

  tail.next = null

  return head
}
