package files

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadFile parses the file and creates
// a map of words to their indices
func ReadFile(path string) (map[string][]int, error) {
	bs, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	m := map[string][]int{}

	for i, w := range strings.Fields(string(bs)) {
		m[w] = append(m[w], i+1)
	}

	return m, nil
}

// absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// minimum
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// O(m + n) find the element in idx1 and the element in idx2
// that are closest to each other and return the distance
func closestPair(idx1, idx2 []int) int {
	if len(idx1) == 1 && len(idx2) == 1 {
		return abs(idx1[0] - idx2[0])
	}

	if len(idx1) == 1 {
		distance := 0
		currentDistance := 0
		i := 0

		for i < len(idx2) && (currentDistance == 0 || distance < currentDistance) {
			currentDistance = distance
			distance = abs(idx1[0] - idx2[i])
			i++
		}

		return min(distance, currentDistance)
	} else if len(idx2) == 1 {
		distance := 0
		currentDistance := 0
		i := 0

		for i < len(idx1) && (currentDistance == 0 || distance < currentDistance) {
			currentDistance = distance
			distance = abs(idx2[0] - idx1[i])
			i++
		}

		return min(distance, currentDistance)
	}

	i := 0
	j := 0
	minDistance := -1
	distance := 0

	for i < len(idx1)-1 && j < len(idx2)-1 {
		distance = abs(idx1[i] - idx2[j])
		iDistance := abs(idx1[i+1] - idx2[j])
		jDistance := abs(idx1[i] - idx2[j+1])

		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}

		if iDistance < jDistance {
			i++
		} else {
			j++
		}
	}

	currentDistance := distance

	if i == len(idx1)-1 {
		for j < len(idx2) && distance <= currentDistance {
			j++
			currentDistance = distance
			distance = abs(idx1[i] - idx2[j])
		}
	} else if i != len(idx1)-1 {
		for i < len(idx1) && distance <= currentDistance {
			i++
			currentDistance = distance
			distance = abs(idx1[i] - idx2[j])
		}
	}

	return min(minDistance, currentDistance)
}

// MinWordDistance finds the minimum distance (number of words)
// between two words in a text file
func MinWordDistance(m map[string][]int, w1, w2 string) (int, error) {
	idx1, ok1 := m[w1]
	idx2, ok2 := m[w2]

	if !ok1 && !ok2 {
		return 0, fmt.Errorf("neither %s or %s appear in the file", w1, w2)
	}
	if !ok1 {
		return 0, fmt.Errorf("%s does not appear in the file", w1)
	}
	if !ok2 {
		return 0, fmt.Errorf("%s does not appear in the file", w2)
	}

	return closestPair(idx1, idx2), nil
}
