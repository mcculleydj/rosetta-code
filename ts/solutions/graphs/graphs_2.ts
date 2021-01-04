import { Vertex } from './graphs_1'

interface Dependency {
  blocker: string
  project: string
}

interface GraphData {
  projectVertices: { [key: string]: Vertex }
  roots: Set<Vertex>
  dependencyMap: { [key: string]: Vertex[] }
}

// O(p + d)
function buildGraph(projects: string[], dependencies: Dependency[]): GraphData {
  const projectVertices: { [key: string]: Vertex } = {}
  const roots: Set<Vertex> = new Set()
  const dependencyMap: { [key: string]: Vertex[] } = {}

  projects.forEach(p => {
    const vertex = new Vertex(p)
    projectVertices[p] = vertex
    roots.add(vertex)
  })

  dependencies.forEach(({ blocker, project }) => {
    roots.delete(projectVertices[project])

    projectVertices[blocker].adjacency.push(projectVertices[project])

    if (dependencyMap[project]) {
      dependencyMap[project].push(projectVertices[blocker])
    } else {
      dependencyMap[project] = [projectVertices[blocker]]
    }
  })

  return {
    projectVertices,
    roots,
    dependencyMap,
  }
}

function buildSchedule(
  projects: string[],
  dependencies: Dependency[],
): string[] {
  const { projectVertices, roots, dependencyMap } = buildGraph(
    projects,
    dependencies,
  )

  const schedule: string[] = []
  const visited: Set<Vertex> = new Set()
  const numVertices = Object.keys(projectVertices).length

  for (const r of roots) {
    schedule.push(...walk(projectVertices[r.data], visited, dependencyMap))

    if (schedule.length === numVertices) {
      break
    }
  }

  if (schedule.length !== numVertices) {
    throw 'no viable path'
  }

  return schedule
}

function walk(
  vertex: Vertex,
  visited: Set<Vertex>,
  dependencyMap: { [key: string]: Vertex[] },
): string[] {
  if (
    dependencyMap[vertex.data] &&
    !dependencyMap[vertex.data].every(v => visited.has(v))
  ) {
    return []
  }

  visited.add(vertex)
  // '' + is a type assertion so that
  // we don't have to define Vertex more than once
  // generally this is not necessary
  const schedule: string[] = ['' + vertex.data]

  for (const v of vertex.adjacency) {
    if (!visited.has(v)) {
      schedule.push(...walk(v, visited, dependencyMap))
    }
  }

  return schedule
}
