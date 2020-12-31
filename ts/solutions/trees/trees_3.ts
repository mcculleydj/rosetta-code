import { TreeNode } from './trees_1'

// for TreeNode class see Minimal BST

interface NodeState {
  isBalanced: boolean
  height: number
}

// O(n)
function isBalanced(tn: TreeNode): NodeState {
  if (!tn) {
    return { isBalanced: true, height: -1 }
  }

  const left = isBalanced(tn.left)
  if (!left.isBalanced) return left

  const right = isBalanced(tn.right)
  if (!right.isBalanced) return right

  if (Math.abs(left.height - right.height) > 1) {
    return { isBalanced: false, height: 0 }
  }

  return {
    isBalanced: true,
    height: 1 + Math.max(left.height, right.height),
  }
}

function isBalancedWrapper(tn: TreeNode): boolean {
  return isBalanced(tn).isBalanced
}
