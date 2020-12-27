// O(n)
// could rely on built in replace method, but that seems to be what the problem is asking us to implement
function URLEncode(s: string): string {
  const charArray: string[] = []
  for (const c of s) {
    if (c === ' ') {
      charArray.push('%20')
    } else {
      charArray.push(c)
    }
  }
  return charArray.join('')
}
