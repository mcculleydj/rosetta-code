// O(n + m)
function checkPermutation(s1: string, s2: string): boolean {
  if (s1.length !== s2.length) {
    return false
  }

  const charCount:{[key: string]: number} = {}

  for (let c of s1) {
    if (charCount[c]) {
      charCount[c]++
    } else {
      charCount[c] = 1
    }
  }

  for (let c of s2) {
    if (!charCount[c]) {
      return false
    }
    charCount[c]--
  }

  return true
}
