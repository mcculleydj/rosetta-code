import { TreeNode } from './trees_1'

// for TreeNode class see Minimal BST

// O(n)
function isBST(
  tn: TreeNode,
  leftBound: number = null,
  rightBound: number = null,
): boolean {
  if (!tn) return true

  if (leftBound !== null && tn.data <= leftBound) {
    return false
  }

  if (rightBound !== null && tn.data > rightBound) {
    return false
  }

  return (
    isBST(tn.left, leftBound, tn.data) && isBST(tn.right, tn.data, rightBound)
  )
}
