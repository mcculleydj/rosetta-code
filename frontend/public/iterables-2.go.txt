package iterables

// CheckPermutation runs in O(n + m)
func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := map[rune]int{}

	for _, c := range s1 {
		if _, ok := charCount[c]; ok {
			charCount[c]++
		} else {
			charCount[c] = 1
		}
	}

	for _, c := range s2 {
		if n, ok := charCount[c]; !ok || n == 0 {
			return false
		}
		charCount[c]--
	}

	return true
}
