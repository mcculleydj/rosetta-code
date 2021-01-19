import { TreeNode } from './trees_1'

// for TreeNode class see Minimal BST

// O(1)
function updatePaths(
  paths: { [key: number]: number },
  key: number,
  delta: number,
) {
  const updatedValue = (paths[key] || 0) + delta

  if (updatedValue === 0) {
    delete paths[key]
  } else {
    paths[key] = updatedValue
  }
}

// O(n)
function pathsToSum(
  node: TreeNode,
  target: number,
  current: number,
  paths: { [key: number]: number },
): number {
  if (!node) {
    return 0
  }

  current += node.data
  let total = paths[current - target] || 0

  if (current === target) {
    total++
  }

  updatePaths(paths, current, 1)
  total += pathsToSum(node.left, target, current, paths)
  total += pathsToSum(node.right, target, current, paths)
  updatePaths(paths, current, -1)

  return total
}
