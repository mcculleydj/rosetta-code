import { TreeNode } from './trees_1'

// see Minimal BST for TreeNode class def

function constructDepthMap(
  tn: TreeNode,
  depth: number,
  depthMap: { [key: number]: number[] },
): { [key: number]: number[] } {
  if (!tn) return

  if (depthMap[depth]) {
    depthMap[depth].push(tn.data)
  } else {
    depthMap[depth] = [tn.data]
  }

  constructDepthMap(tn.left, depth + 1, depthMap)
  constructDepthMap(tn.right, depth + 1, depthMap)

  return depthMap
}

function swap(ns: number[], i: number, j: number) {
  const temp = ns[i]
  ns[i] = ns[j]
  ns[j] = temp
}

function permutations(ns: number[], k: number, acc: number[][]) {
  if (k === ns.length) {
    acc.push([...ns])
  } else {
    for (let i = k; i < ns.length; i++) {
      swap(ns, i, k)
      permutations(ns, k + 1, acc)
      swap(ns, i, k)
    }
  }
}

function bstSequences(tn: TreeNode): number[][] {
  const depthMap = constructDepthMap(tn, 0, {})
  let sequences: number[][] = [[tn.data]]

  for (let d = 1; d < Object.keys(depthMap).length; d++) {
    const appendedSequences: number[][] = []
    const ps: number[][] = []
    permutations(depthMap[d], 0, ps)
    for (const p of ps) {
      for (const s of sequences) {
        appendedSequences.push(s.concat(p))
      }
    }
    sequences = appendedSequences
  }

  return sequences
}
