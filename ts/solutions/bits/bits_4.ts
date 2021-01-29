// O(n) -- as a reminder bitwise ops cast to 32-bit in JS
function flipBit(n: number): number {
  let i = 0
  const zeroIndices: number[] = []

  while (n > 0) {
    if ((n & 1) === 0) {
      zeroIndices.push(i)
    }
    n >>= 1
    i++
  }

  if (zeroIndices.length === 0) {
    return i - 1
  } else if (zeroIndices.length === 1) {
    return i
  }

  let longest = zeroIndices[1]

  zeroIndices.forEach((idx, j) => {
    if (zeroIndices.length - 1 === j) return

    const start = idx + 1
    let end: number

    if (zeroIndices.length - 2 === j) {
      end = i
    } else {
      end = zeroIndices[j + 2]
    }

    if (end - start > longest) {
      longest = end - start
    }
  })

  return longest
}
