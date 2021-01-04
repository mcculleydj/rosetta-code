class Node:
    def __init__(self, data=None):
        self.data = data
        self.neighbors = []

    def __repr__(self):
        return f'Data: {self.data} | Edges: {[neighbor.data for neighbor in self.neighbors]}'


# O(n) in both space and time
def path_exists(source, target):
    visited = set()
    dfs(source, target, visited)

    return target in visited


# basic DFS implementation
def dfs(node, target, visited=set()):
    visited.add(node)

    for neighbor in node.neighbors:
        # stop any additional recursion once the target is discovered
        if target in visited:
            break
        if neighbor not in visited:
            dfs(neighbor, target, visited)
