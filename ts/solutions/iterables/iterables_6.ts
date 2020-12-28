// O(n)
function compressString(s: string): string {
  if (s === '') {
    return ''
  }

  const acc: string[] = []
  let startIndex = 0
  let currentIndex = 0
  let currentChar = s[0]

  for (const c of s) {
    if (c !== currentChar) {
      acc.push(currentChar, '' + (currentIndex - startIndex))
      startIndex = currentIndex
      currentChar = c
    }
    currentIndex++
  }

  acc.push(currentChar, '' + (currentIndex - startIndex))

  return acc.length > s.length ? s : acc.join('')
}

console.log(compressString('aabcccccaaa'))
