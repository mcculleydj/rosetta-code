import { LinkedList, Node_ } from './linked_lists_1'

// see Remove Duplicates for List and Node type definitions

// O(n)
function kthToLastElement(l: LinkedList, k: number): Node_ {
  let leadNode: Node_ = l.head
  let lagNode: Node_ = l.head
  let leadIndex = 0
  let lagIndex = 0

  while (leadNode) {
    leadNode = leadNode.next
    leadIndex++
    if (leadIndex > k) {
      lagNode = lagNode.next
      lagIndex++
    }
  }

  if (leadIndex - lagIndex < k) {
    console.log(
      `List is not long enough to have a ${k}(st/nd/rd/th) to last element, returning the first node instead.`,
    )
  }

  return lagNode
}
