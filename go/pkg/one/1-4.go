package one

// PalindromePermutation runs in O(n)
func PalindromePermutation(s string) bool {
	m := map[rune]int{}

	for _, c := range s {
		// in Go, there is an implied zero value based on type
		// therefore, you don't need the existence check common in
		// in other languages before incrementing the count
		m[c]++
	}

	hasOdd := false

	for k := range m {
		if m[k]%2 == 1 {
			if hasOdd {
				return false
			}
			hasOdd = true
		}
	}

	return true
}
