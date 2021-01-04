import { TreeNode } from './trees_1'

// for TreeNode class see Minimal BST

class QueueNode {
  treeNode: TreeNode
  depth: number
  next: QueueNode

  constructor(treeNode?: TreeNode, depth?: number, next?: QueueNode) {
    this.treeNode = treeNode
    this.depth = depth
    this.next = next
  }

  toString(): string {
    return `(${this.treeNode.data}, ${this.depth})`
  }
}

class Queue {
  head: QueueNode
  tail: QueueNode
  length: number

  constructor(head?: QueueNode, tail?: QueueNode) {
    this.head = head
    this.tail = tail
    this.length = 0
  }

  toString(): string {
    const ss: string[] = []
    let node = this.head

    while (node) {
      ss.push(node.toString())
      node = node.next
    }

    return ss.join(' -> ')
  }

  add(treeNode: TreeNode, depth: number): void {
    const node = new QueueNode(treeNode, depth)

    if (!this.head) {
      this.head = node
    }

    if (this.tail) {
      this.tail.next = node
    }

    this.tail = node
    this.length++
  }

  remove(): QueueNode {
    if (!this.head) {
      throw 'queue is empty'
    }

    const node = this.head
    this.head = this.head.next

    if (this.length === 1) {
      this.tail = null
    }

    this.length--

    return node
  }
}

// O(n)
function listDepths(
  treeNode: TreeNode,
  depth = 0,
  queue: Queue = new Queue(),
  depths: Queue[] = [],
): Queue[] {
  if (depth == depths.length) {
    depths.push(new Queue())
  }

  depths[depth].add(treeNode, depth)

  if (treeNode.left) {
    queue.add(treeNode.left, depth + 1)
  }
  if (treeNode.right) {
    queue.add(treeNode.right, depth + 1)
  }

  if (queue.length == 0) {
    return depths
  }

  const { treeNode: nextNode, depth: nextDepth } = queue.remove()
  return listDepths(nextNode, nextDepth, queue, depths)
}
