// Node already exists in JS
export class Node_ {
  value: number
  next: Node_

  constructor(value: number, next: Node_) {
    this.value = value
    this.next = next
  }
}

export class LinkedList {
  head: Node_

  constructor(head: Node_) {
    this.head = head
  }

  toString() {
    let s = ''
    let node = this.head
    while (node) {
      s = `${s}${node.value} -> `
      node = node.next
    }
    return s + 'nil'
  }
}

// O(n)
function removeDuplicatesTime(l: LinkedList) {
  const valueSet: Set<number> = new Set()
  let node = l.head

  while (node.next) {
    if (valueSet.has(node.next.value)) {
      node.next = node.next.next
    } else {
      valueSet.add(node.next.value)
      node = node.next
    }
  }
}

// O(1) space -- O(n^2) time
function removeDuplicatesSpace(l: LinkedList) {
  let indexNode = l.head
  let cursorNode

  while (indexNode.next) {
    cursorNode = indexNode
    while (cursorNode.next) {
      if (cursorNode.next.value === indexNode.value) {
        cursorNode.next = cursorNode.next.next
      } else {
        cursorNode = cursorNode.next
      }
    }
    indexNode = indexNode.next
  }
}
