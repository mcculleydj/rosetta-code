package iterables

import "strings"

// URLEncode replaces all spaces with %20
//could rely on built in replace method, but that seems to be what the problem is asking us to implement
func URLEncode(s string) string {
	// amortized cost to append to a slice in Go is O(1) even if O(n) in the worst case
	ss := []string{}

	// O(n) where n is the length of the string
	for _, c := range s {
		if c == ' ' {
			ss = append(ss, "%20")
		} else {
			// complexity to cast is O(m) where m is the number of bytes
			// for each rune comprising the string this is effectively O(1)
			ss = append(ss, string(c))
		}
	}

	// strings.Join is O(n)
	return strings.Join(ss, "")
}
