function getMSB(n: number): number {
  if (n === 0) return 0
  let msb = 0

  for (let i = 0; i < 32; i++) {
    const b = n & 1
    n >>= 1

    if (b === 1) {
      msb = i
    }
  }

  return msb
}

function insert(m: number, n: number, i: number): number {
  const msb = getMSB(m)
  m <<= i
  const left = -1 << (i + msb)
  const right = (1 << i) - 1
  const mask = left | right
  n &= mask

  return m | n
}
