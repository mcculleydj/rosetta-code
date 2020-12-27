// O(n)
function palindromePermutation(s: string): boolean {
  const charCount: { [key: string]: number } = {}
  for (const c of s) {
    if (charCount[c]) {
      charCount[c]++
    } else {
      charCount[c] = 1
    }
  }

  let hasOdd = false

  // Object.values requires a later TS transpilation target than ES6
  for (const char in charCount) {
    if (charCount[char] % 2 === 1) {
      if (hasOdd) {
        return false
      }
      hasOdd = true
    }
  }

  return true
}
