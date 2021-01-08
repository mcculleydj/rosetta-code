class BSTNode {
  data: number
  left: BSTNode
  right: BSTNode
  leftSize: number
  rightSize: number

  constructor(data: number) {
    this.data = data
    this.left = null
    this.right = null
    this.leftSize = 0
    this.rightSize = 0
  }

  toString(): string {
    return `${this.data} | ${this.leftSize} | ${this.rightSize} => (${this.left}, ${this.right})`
  }
}

interface FindReturn {
  node: BSTNode
  parent: BSTNode
  direction: string
}

class BST {
  root: BSTNode

  constructor(root?: BSTNode) {
    this.root = root
  }

  toString(): string {
    return this.root.toString()
  }

  insert(data: number): void {
    if (!this.root) {
      this.root = new BSTNode(data)
      return
    }

    let node = this.root
    let parent: BSTNode
    let direction: string

    while (node) {
      parent = node

      if (data <= node.data) {
        parent.leftSize++
        node = node.left
        direction = 'left'
      } else {
        parent.rightSize++
        node = node.right
        direction = 'right'
      }
    }

    if (direction === 'left') {
      parent.left = new BSTNode(data)
    } else {
      parent.right = new BSTNode(data)
    }
  }

  find(data: number, decrement = 0): FindReturn {
    if (!this.root) {
      return {
        node: null,
        parent: null,
        direction: '',
      }
    }

    if (this.root.data === data) {
      return {
        node: this.root,
        parent: null,
        direction: '',
      }
    }

    let node = this.root
    let parent: BSTNode
    let direction: string
    const leftDecrementedNodes: BSTNode[] = []
    const rightDecrementedNodes: BSTNode[] = []

    while (node && node.data !== data) {
      parent = node

      if (data <= node.data) {
        parent.leftSize -= decrement
        leftDecrementedNodes.push(parent)
        node = node.left
        direction = 'left'
      } else {
        parent.rightSize -= decrement
        rightDecrementedNodes.push(parent)
        node = node.right
        direction = 'right'
      }
    }

    if (!node && decrement > 0) {
      for (const decrementedNode of leftDecrementedNodes) {
        decrementedNode.leftSize++
      }

      for (const decrementedNode of rightDecrementedNodes) {
        decrementedNode.rightSize++
      }
    }

    return { node, parent, direction }
  }

  delete(data: number): void {
    const { node, parent, direction } = this.find(data, 1)

    if (!node) return

    let nodeToShift: BSTNode

    if (node.left && node.right) {
      nodeToShift = node.left
      let parentOfNodeToShift: BSTNode

      while (nodeToShift.right) {
        parentOfNodeToShift = nodeToShift
        parentOfNodeToShift.rightSize--
        nodeToShift = nodeToShift.right
      }

      parentOfNodeToShift.right = nodeToShift.left
      nodeToShift.leftSize = node.leftSize - 1
      nodeToShift.rightSize = node.rightSize
    } else if (node.left) {
      nodeToShift = node.left
      nodeToShift.leftSize = node.leftSize - 1
    } else if (node.right) {
      nodeToShift = node.right
      nodeToShift.rightSize = node.rightSize - 1
    }

    if (parent) {
      if (direction === 'left') {
        parent.left = nodeToShift
      } else {
        parent.right = nodeToShift
      }
    } else {
      this.root = nodeToShift
    }

    if (nodeToShift) {
      if (node.left != nodeToShift) {
        nodeToShift.left = node.left
      }
      if (node.right != nodeToShift) {
        nodeToShift.right = node.right
      }
    }
  }

  random(): BSTNode {
    if (!this.root) return null

    let node = this.root

    while (node) {
      const n = Math.random()
      const total = 1 + node.leftSize + node.rightSize

      if (n <= 1 / total) {
        return node
      } else if (n <= (1 + node.leftSize) / total) {
        node = node.left
      } else {
        node = node.right
      }
    }

    return null
  }
}
