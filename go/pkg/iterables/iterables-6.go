package iterables

import (
	"strconv"
	"strings"
)

// CompressString runs in O(n)
func CompressString(s string) string {
	// handle empty string
	if s == "" {
		return ""
	}

	ss := []string{}
	startIndex := 0
	currentIndex := 0
	currentChar := rune(s[0])

	// O(n)
	for _, c := range s {
		if c != currentChar {
			ss = append(ss, []string{string(currentChar), strconv.Itoa(currentIndex - startIndex)}...)
			startIndex = currentIndex
			currentChar = c
		}
		currentIndex++
	}

	// handle the final character sequence
	ss = append(ss, []string{string(currentChar), strconv.Itoa(currentIndex - startIndex)}...)

	if len(ss) < len(s) {
		// O(n)
		return strings.Join(ss, "")
	}

	return s
}
