class BTNode {
  data: number
  left: BTNode
  right: BTNode

  constructor(data?: number, left?: BTNode, right?: BTNode) {
    this.data = data
    this.left = left
    this.right = right
  }

  toString(): string {
    if (!this.left && !this.right) return '' + this.data
    else if (!this.left) return `${this.data} (null, ${this.right.toString()})`
    else if (!this.right) return `${this.data} (${this.left.toString()}, null)`
    return `${this.data} (${this.left.toString()}, ${this.right.toString()})`
  }
}

interface SplitList {
  middle: number
  left: number[]
  right: number[]
}

// O(1)
function splitList(ns: number[]): SplitList {
  const i = Math.floor(ns.length / 2)
  return {
    middle: ns[i],
    left: ns.slice(0, i),
    right: ns.slice(i + 1),
  }
}

// O(n)
function minBST(ns: number[], node?: BTNode) {
  if (!node) {
    node = new BTNode()
  }

  const { middle, left, right } = splitList(ns)
  node.data = middle

  if (left.length) {
    node.left = new BTNode()
    minBST(left, node.left)
  }

  if (right.length) {
    node.right = new BTNode()
    minBST(right, node.right)
  }

  return node
}
