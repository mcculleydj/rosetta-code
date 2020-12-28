// O(n)
function isRotation(s1: string, s2: string): boolean {
  if (s1.length !== s2.length) {
    return false
  }

  return (s2 + s2).includes(s1)
}
