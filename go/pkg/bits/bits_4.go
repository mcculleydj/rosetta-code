package bits

// FlipBit is O(n) -- n is the number of bits bounded by 64 in this case
func FlipBit(n int64) int {
	i := 0
	zeroIndices := []int{}

	for n > 0 {
		if n&1 == 0 {
			zeroIndices = append(zeroIndices, i)
		}
		n = n >> 1
		i++
	}

	if len(zeroIndices) == 0 {
		return i - 1
	} else if len(zeroIndices) == 1 {
		return i
	}

	longest := zeroIndices[1]

	for j, idx := range zeroIndices {
		if len(zeroIndices)-1 == j {
			break
		}

		start := idx + 1
		var end int

		if len(zeroIndices)-2 == j {
			end = i
		} else {
			end = zeroIndices[j+2]
		}

		if end-start > longest {
			longest = end - start
		}
	}

	return longest
}
