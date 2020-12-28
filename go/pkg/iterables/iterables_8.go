package iterables

// ZeroMatrix runs in O(nm)
func ZeroMatrix(m [][]int) {
	zeroRows := map[int]bool{}
	zeroColumns := map[int]bool{}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				zeroRows[i] = true
				zeroColumns[j] = true
			}
		}
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if zeroRows[i] || zeroRows[j] {
				m[i][j] = 0
			}
		}
	}
}
