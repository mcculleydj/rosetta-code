package bits

import (
	"errors"
)

// Add is O(n) where n is the number of bits necessary
// to represent the larger number
func Add(x, y int64) (int64, error) {
	for y != 0 {
		// produce a bit vector where both bits are 1s
		// and we will need to carry over to the left
		carry := x & y

		// XOR is the result of addition without carrying
		x = x ^ y

		// store value of y before the shift
		y0 := y

		// shift the carry vector and repeat this process
		y = carry << 1

		// if y ever shifts from positive to negative or vice versa
		// the sum of the two numbers has overflowed the int64
		if (y0 < 0 && y > 0) || (y0 > 0 && y < 0) {
			return 0, errors.New("int64 overflow")
		}
	}

	return x, nil
}
