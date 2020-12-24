package one

// IsUnique with access to O(1) lookup data structures O(n)
func IsUnique(s string) bool {
	m := map[rune]bool{}
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}

// IsUniqueSlow without access to any additional data structures O(n^2)
// could sort in O(n log(n)), but the assumption is that a new string
// constitutes an additional data structure
func IsUniqueSlow(s string) bool {
	for i, c := range s {
		for j := i + 1; j < len(s); j++ {
			if c == rune(s[j]) {
				return false
			}
		}
	}
	return true
}
