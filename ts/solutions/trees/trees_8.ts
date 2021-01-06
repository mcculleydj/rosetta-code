import { TreeNode } from './trees_1'

// see Minimal BST for TreeNode class def

// O(n)
function traverseTree(
  tn: TreeNode,
  depth: number,
  heightToNodes: { [key: number]: TreeNode[] },
  target: number,
): number {
  if (!tn) {
    return depth - 1
  }

  const leftHeight = traverseTree(tn.left, depth + 1, heightToNodes, target)
  const rightHeight = traverseTree(tn.right, depth + 1, heightToNodes, target)

  const height = Math.max(leftHeight, rightHeight)

  // assumes no tn.data == null
  // tn itself is null in this case
  if (tn.data === target) {
    if (heightToNodes[height - depth]) {
      heightToNodes[height - depth].push(tn)
    } else {
      heightToNodes[height - depth] = [tn]
    }
  }

  return height
}

// O(n)
function valueEquivalent(t1: TreeNode, t2: TreeNode): boolean {
  if (!t1 && !t2) {
    return true
  } else if (t1.data !== t2.data) {
    return false
  }

  return (
    valueEquivalent(t1.left, t2.left) && valueEquivalent(t1.right, t2.right)
  )
}

// O(m + n)
function checkSubtree(t1: TreeNode, t2: TreeNode): boolean {
  const height = traverseTree(t2, 0, null, null)
  const heightToNodes: { [key: number]: TreeNode[] } = {}
  traverseTree(t1, 0, heightToNodes, t2.data)

  for (const node of heightToNodes[height]) {
    if (valueEquivalent(node, t2)) {
      return true
    }
  }

  return false
}
