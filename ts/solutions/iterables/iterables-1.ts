// with access to O(1) lookup data structures O(n)
function isUnique(s: string): boolean {
  const charSet: Set<string> = new Set()
  for (const c of s) {
    if (charSet.has(c)) {
      return false
    }
    charSet.add(c)
  }
  return true
}

// without access to any additional data structures O(n^2)
// could sort in O(n log(n)), but the assumption is that a new string
// constitutes an additional data structure
function isUniqueSlow(s: string): boolean {
  let i = 0
  for (const c of s) {
    for (let j = i + 1; j < s.length; j++) {
      if (c === s[j]) {
        return false
      }
    }
    i++
  }
  return true
}
