package bits

func getMSB(n int32) int32 {
	if n == 0 {
		return 0
	}

	var msb int32

	for i := 0; i < 32; i++ {
		b := n & 1
		n = n >> 1

		if b == 1 {
			msb = int32(i)
		}
	}

	return msb
}

// Insert splices m into n
func Insert(m, n, i int32) int32 {
	var left, right int32
	msb := getMSB(m)
	m = m << i
	left = -1 << (i + msb)
	right = (1 << i) - 1
	mask := left | right
	n = n & mask

	return m | n
}
