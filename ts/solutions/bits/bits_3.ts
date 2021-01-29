// O(32)
function binaryToString(n: number): string {
  const bits: string[] = []

  while (n > 0) {
    if (bits.length > 31) throw 'loss of precision'

    const k = n * 2

    if (k >= 1) {
      bits.push('1')
      n = k - 1
    } else {
      bits.push('0')
      n = k
    }
  }

  return '.' + bits.join('')
}
