from graphs_1 import Node

# for Node class def see Path Existence


# O(m) where m is the number of dependencies, which is O(n^2) where n is the number of nodes
# create each Node instance defining the graph
# return a set of roots,, where a root is a Node without any dependencies
# dependencies: [(A, B), (C, D), ...] => B depends on A, D depends on C, and so on
# for simplicity assume that a tuple is not repeated in this set
# projects = [A, B, C, D, ...]
def build_graph(projects, dependencies):
    project_nodes = {}

    # O(n)
    # will use this to figure out which projects have no dependencies
    roots = set(projects)

    # for every node we want a set of dependencies
    # to know if the visitation constraint will be satisfied
    dependency_map = {}

    # O(n)
    # ensure that project_nodes represents the complete set
    for project in projects:
        project_nodes[project] = Node(project)

    # O(m)
    for dependency, project in dependencies:
        if project in roots:
            roots.remove(project)

        # add directed edge
        project_nodes[dependency].neighbors.append(project_nodes[project])

        # populate dependency map
        if dependency_map.get(project):
            dependency_map[project].add(project_nodes[dependency])
        else:
            dependency_map[project] = {project_nodes[dependency]}

    return project_nodes, roots, dependency_map


# O(m + n) || O(n^2) -- see comment above
def build_schedule(projects, dependencies):
    project_nodes, roots, dependency_map = build_graph(projects, dependencies)

    # would use a linked list for O(1) list extension
    # using a built-in list for convenience
    schedule = []
    visited = set()

    for root in roots:
        schedule += walk(project_nodes[root], visited, dependency_map)

        # viable schedule discovered
        if len(schedule) == len(project_nodes):
            break

    # cycle detected
    if len(schedule) != len(project_nodes):
        raise Exception('no viable build order')

    return schedule


# O(n) DFS with the addition of a dependency_map to ensure that a node is visitable
def walk(node, visited, dependency_map):
    if dependency_map.get(node.data):
        # builds a boolean vector in O(n) and then folds it in O(n)
        if not all([dependency in visited for dependency in dependency_map.get(node.data)]):
            return []

    visited.add(node)
    schedule = [node.data]

    # only touching each node once O(n)
    for neighbor in node.neighbors:
        if neighbor not in visited:
            # technically list concatnation as the stack unwinds is O(n^2)
            # but, this could've been implemented using linked lists
            # since we have no need for random index access
            # which would keep the time complexity at O(n)
            schedule += walk(neighbor, visited, dependency_map)

    return schedule
