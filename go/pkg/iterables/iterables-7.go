package iterables

// RotateMatrix runs in O(n^2)
func RotateMatrix(m [][]int) {
	for i := 0; i < len(m)/2; i++ {
		for j := i; j < len(m)-i-1; j++ {
			// slices are pointers to underlying arrays
			// so while everything is Go is pass by value, this is already
			// a reference and the modifications will occur in place
			rotateCell(m, m[i][j], [2]int{j, len(m) - i - 1}, [2]int{i, j})
		}
	}
}

func rotateCell(m [][]int, value int, destination, start [2]int) {
	previousValue := m[destination[0]][destination[1]]
	m[destination[0]][destination[1]] = value

	// base case
	if destination[0] == start[0] && destination[1] == start[1] {
		return
	}

	rotateCell(m, previousValue, [2]int{destination[1], len(m) - destination[0] - 1}, start)
}
