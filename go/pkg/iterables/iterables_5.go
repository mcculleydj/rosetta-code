package iterables

// O(n)
func hasAdditionalCharacter(s1, s2 string) bool {
	hasSkipped := false
	i := 0
	j := 0

	for i < len(s1) {
		if s1[i] == s2[j] {
			i++
			j++
		} else if !hasSkipped {
			hasSkipped = true
			j++
		} else {
			return false
		}
	}

	return true
}

// OneAway runs in O(n)
func OneAway(s1, s2 string) bool {
	if len(s2)-len(s1) == 1 {
		return hasAdditionalCharacter(s1, s2)
	} else if len(s1)-len(s2) == 1 {
		return hasAdditionalCharacter(s2, s1)
	} else if len(s1) == len(s2) {
		seenReplacement := false
		for i, c := range s1 {
			if c != rune(s2[i]) {
				if seenReplacement {
					return false
				}
				seenReplacement = true
			}
		}
		return seenReplacement
	}
	return false
}
