import { LinkedList, Node_ } from './linked_1'

// see Remove Duplicates for List and Node type definitions

// O(m + n)
function findIntersection(l1: LinkedList, l2: LinkedList): Node_ {
  const nodeSet: Set<Node_> = new Set()

  let node = l1.head
  while (node) {
    nodeSet.add(node)
    node = node.next
  }

  node = l2.head
  while (node) {
    if (nodeSet.has(node)) {
      return node
    }
    node = node.next
  }

  return null
}
