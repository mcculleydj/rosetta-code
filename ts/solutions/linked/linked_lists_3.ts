import { Node_ } from './linked_lists_1'

// O(1) -- see Python for comments
function deleteNode(n: Node_) {
  n.value = n.next.value
  n.next = n.next.next
}
