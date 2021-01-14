package files

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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

func binarySearch(n int, idx []int) (int, error) {
	if len(idx) == 0 {
		return 0, errors.New("index slice is empty")
	}

	if n < idx[0] {
		return idx[0] - n, nil
	}
	if n > idx[len(idx)-1] {
		return n - idx[len(idx)-1], nil
	}
	if len(idx) == 2 {
		return int(math.Min(float64(n-idx[0]), float64(idx[1]-n))), nil
	}

	i := 0
	j := len(idx) - 1
	span := j - i

	for span > 1 {
		middle := idx[i+span/2]

		if n < middle {
			j = i + span/2
		} else {
			i = i + span/2
		}

		span = j - i
	}

	return int(math.Min(float64(n-idx[i]), float64(idx[j]-n))), nil
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

	var outer, inner []int

	if len(idx1) > len(idx2) {
		outer = idx2
		inner = idx1
	} else {
		outer = idx1
		inner = idx2
	}

	minDistance := -1

	for _, n := range outer {
		min, err := binarySearch(n, inner)
		if err != nil {
			return 0, err
		}

		if minDistance == -1 || min < minDistance {
			minDistance = min
		}
	}

	return minDistance, nil
}
