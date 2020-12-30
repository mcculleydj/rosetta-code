package graphs

import (
	"fmt"
	"strconv"
	"strings"
)

// Node is a member of a graph defined by the data it holds
// and the links to other nodes
type Node struct {
	Data int
	Adj  []*Node
}

// ToString produces a string representing a node and its adjacencies
func (n *Node) ToString() string {
	ss := []string{}
	for _, node := range n.Adj {
		ss = append(ss, strconv.Itoa(node.Data))
	}
	return fmt.Sprintf("Data: %d | Edges: %s", n.Data, strings.Join(ss, ", "))
}

// PathExists determines if a path exists between two nodes in a graph
// in O(n) time and space
func PathExists(source, target *Node) bool {
	visited := map[*Node]bool{}
	dfs(source, target, visited)

	return visited[target]
}

// basic DFS implementation
func dfs(node, target *Node, visited map[*Node]bool) {
	visited[node] = true
	for _, adj := range node.Adj {
		if visited[target] {
			break
		}
		if !visited[adj] {
			dfs(adj, target, visited)
		}
	}
}
