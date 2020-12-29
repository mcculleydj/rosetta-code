import { LinkedList, Node_ } from './linked_1'

// see Remove Duplicates for List and Node type definitions

// O(n)
function listToStack(l: LinkedList): Node_[] {
  const stack: Node_[] = []
  let node = l.head
  while (node) {
    stack.push(node)
    node = node.next
  }
  return stack
}

// O(n)
function isPalindrome(l: LinkedList): boolean {
  let listNode: Node_ = l.head
  const stack = listToStack(l)
  const listLength = stack.length

  while (stack.length >= Math.floor(listLength / 2)) {
    const stackNode = stack.pop()
    if (stackNode.value !== listNode.value) {
      return false
    }
    listNode = listNode.next
  }

  return true
}
