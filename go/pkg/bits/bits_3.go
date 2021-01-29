package bits

import (
	"errors"
	"strings"
)

// BinaryToString performs at most 32 iterations
func BinaryToString(n float64) (string, error) {
	bits := []string{}
	for n > 0 {
		if len(bits) > 31 {
			return "", errors.New("loss of precision")
		}

		k := n * 2

		if k >= 1 {
			bits = append(bits, "1")
			n = k - 1
		} else {
			bits = append(bits, "0")
			n = k
		}
	}

	return "." + strings.Join(bits, ""), nil
}
