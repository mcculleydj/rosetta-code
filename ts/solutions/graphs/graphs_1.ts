export class Vertex {
  data: number | string
  adjacency: Vertex[]

  constructor(data?: number | string) {
    this.data = data
    this.adjacency = []
  }

  toString(): string {
    return `Data: ${this.data} | Edges: ${this.adjacency.map(n => n.data)}`
  }
}

// O(n) time and space
function pathExists(source: Vertex, target: Vertex): boolean {
  const visited: Set<Vertex> = new Set()
  dfs(source, target, visited)
  return visited.has(target)
}

function dfs(node: Vertex, target: Vertex, visited: Set<Vertex>) {
  visited.add(node)
  for (const adj of node.adjacency) {
    if (visited.has(target)) break
    if (!visited.has(adj)) dfs(adj, target, visited)
  }
}
