export class TreeNodeWithParent {
  data: number
  parent: TreeNodeWithParent
  left: TreeNodeWithParent
  right: TreeNodeWithParent

  constructor(
    data?: number,
    parent?: TreeNodeWithParent,
    left?: TreeNodeWithParent,
    right?: TreeNodeWithParent,
  ) {
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

// O(log n)
function checkRight(tn: TreeNodeWithParent): TreeNodeWithParent {
  while (tn.left) {
    tn = tn.left
  }
  return tn
}

// O(log n)
function checkParent(tn: TreeNodeWithParent): TreeNodeWithParent {
  let parent = tn.parent

  while (parent) {
    if (parent.data >= tn.data) {
      return parent
    }
    parent = parent.parent
  }

  return null
}

// O(log n)
function findNext(tn: TreeNodeWithParent): TreeNodeWithParent {
  if (!tn) return null
  if (tn.right) return checkRight(tn.right)
  return checkParent(tn)
}
