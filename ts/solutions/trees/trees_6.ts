import { TreeNode } from './trees_1'

// see Minimal BST for TreeNode class def

// O(n)
function checkForNodes(
  tn: TreeNode,
  targetOne: TreeNode,
  targetTwo: TreeNode,
): boolean[] {
  if (!tn) return [false, false]

  if (tn === targetOne) {
    const [, hasOther] =
      checkForNodes(tn.left, targetOne, targetTwo) ||
      checkForNodes(tn.right, targetOne, targetTwo)
    return [true, hasOther]
  }

  if (tn === targetTwo) {
    const [hasOther] =
      checkForNodes(tn.left, targetOne, targetTwo) ||
      checkForNodes(tn.right, targetOne, targetTwo)
    return [hasOther, true]
  }

  const [oneInLeft, twoInLeft] = checkForNodes(tn.left, targetOne, targetTwo)
  const [oneInRight, twoInRight] = checkForNodes(tn.right, targetOne, targetTwo)

  return [oneInLeft || oneInRight, twoInLeft || twoInRight]
}

// O(n^2) worst case
// O(n log n) for a balanced tree
function firstCommonAncestor(
  tn: TreeNode,
  targetOne: TreeNode,
  targetTwo: TreeNode,
): TreeNode {
  const [oneInLeft, twoInLeft] = checkForNodes(tn.left, targetOne, targetTwo)
  const [oneInRight, twoInRight] = checkForNodes(tn.right, targetOne, targetTwo)

  if ((oneInLeft && twoInRight) || (oneInRight && twoInLeft)) return tn

  if (oneInLeft && twoInLeft)
    return firstCommonAncestor(tn.left, targetOne, targetTwo)

  return firstCommonAncestor(tn.right, targetOne, targetTwo)
}
