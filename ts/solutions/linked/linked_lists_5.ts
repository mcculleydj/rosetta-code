import { LinkedList, Node_ } from './linked_lists_1'

// see Remove Duplicates for List and Node type definitions

// O(m + n)
function rightToLeftSum(l1: LinkedList, l2: LinkedList): LinkedList {
  let n1: Node_ = l1.head
  let n2: Node_ = l2.head
  let nextValue = n1.value + n2.value
  let sumNode = new Node_(nextValue % 10, null)
  const sumList = new LinkedList(sumNode)
  let carry = Math.floor(nextValue / 10)

  while (n1.next || n2.next) {
    if (!n1.next) {
      nextValue = n2.next.value + carry
      n2 = n2.next
    } else if (!n2.next) {
      nextValue = n1.next.value + carry
      n1 = n1.next
    } else {
      nextValue = n1.next.value + n2.next.value + carry
      n1 = n1.next
      n2 = n2.next
    }

    carry = Math.floor(nextValue / 10)
    sumNode.next = new Node_(nextValue % 10, null)
    sumNode = sumNode.next
  }

  if (carry === 1) {
    sumNode.next = new Node_(1, null)
  }

  return sumList
}

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

// O(m + n)
function leftToRightSum(l1: LinkedList, l2: LinkedList): LinkedList {
  const stack1 = listToStack(l1)
  const stack2 = listToStack(l2)
  let head: Node_
  let node: Node_
  let nextValue
  let carry = 0

  while (stack1.length > 0 || stack2.length > 0) {
    if (stack1.length === 0) {
      node = stack2.pop()
      nextValue = node.value + carry
      node.value = nextValue % 10
      node.next = head
    } else if (stack2.length === 0) {
      node = stack1.pop()
      nextValue = node.value + carry
      node.value = nextValue % 10
      node.next = head
    } else {
      const node1 = stack1.pop()
      const node2 = stack2.pop()
      nextValue = node1.value + node2.value + carry
      node = new Node_(nextValue % 10, head)
    }

    carry = Math.floor(nextValue / 10)
    head = node
  }

  if (carry === 1) {
    node = new Node_(1, head)
    head = node
  }

  return new LinkedList(head)
}
