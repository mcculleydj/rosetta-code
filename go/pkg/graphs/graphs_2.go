package graphs

import (
	"errors"
	"fmt"
	"strings"
)

// StringNode represents a node in the graph
type StringNode struct {
	Data string
	Adj  []*StringNode
}

// ToString produces a string representing a node and its adjacencies
func (n *StringNode) ToString() string {
	ss := []string{}
	for _, node := range n.Adj {
		ss = append(ss, node.Data)
	}
	return fmt.Sprintf("Data: %s | Edges: %s", n.Data, strings.Join(ss, ", "))
}

// Dependency encodes the blocker => project relation
type Dependency struct {
	Blocker string
	Project string
}

// O(p + d) where p is the number of projects and
// d is the number of dependencies
func buildGraph(projects []string, dependencies []Dependency) (map[string]*StringNode, map[string]bool, map[string]map[*StringNode]bool) {
	projectNodes := map[string]*StringNode{}
	roots := map[string]bool{}
	dependencyMap := map[string]map[*StringNode]bool{}

	for _, project := range projects {
		projectNodes[project] = &StringNode{Data: project}
		roots[project] = true
	}

	for _, dependency := range dependencies {
		if roots[dependency.Project] {
			delete(roots, dependency.Project)
		}

		projectNodes[dependency.Blocker].Adj = append(
			projectNodes[dependency.Blocker].Adj,
			projectNodes[dependency.Project],
		)

		if _, ok := dependencyMap[dependency.Project]; !ok {
			dependencyMap[dependency.Project] = map[*StringNode]bool{}
		}

		dependencyMap[dependency.Project][projectNodes[dependency.Blocker]] = true
	}

	return projectNodes, roots, dependencyMap
}

// BuildSchedule runs in O(p + d)
func BuildSchedule(projects []string, dependencies []Dependency) ([]string, error) {
	projectNodes, roots, dependencyMap := buildGraph(projects, dependencies)

	schedule := []string{}
	visited := map[*StringNode]bool{}

	for root := range roots {
		// could be implemented in O(1)
		// using slices out of convenience
		schedule = append(schedule, walk(projectNodes[root], visited, dependencyMap)...)

		if len(schedule) == len(projectNodes) {
			break
		}
	}

	if len(schedule) != len(projectNodes) {
		return nil, errors.New("no viable path")
	}

	return schedule, nil
}

// O(p)
func walk(node *StringNode, visited map[*StringNode]bool, dependencyMap map[string]map[*StringNode]bool) []string {
	if dependencies, ok := dependencyMap[node.Data]; ok {
		for d := range dependencies {
			if !visited[d] {
				return []string{}
			}
		}
	}

	visited[node] = true
	schedule := []string{node.Data}

	for _, adj := range node.Adj {
		if !visited[adj] {
			schedule = append(schedule, walk(adj, visited, dependencyMap)...)
		}
	}

	return schedule
}
