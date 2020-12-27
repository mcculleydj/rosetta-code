// O(n)
function hasAdditionalCharacter(s1: string, s2: string): boolean {
  let hasSkipped = false
  let i = 0
  let j = 0

  while (i < s1.length) {
    if (s1[i] == s2[j]) {
      i++
      j++
    } else if (!hasSkipped) {
      hasSkipped = true
      j++
    } else {
      return false
    }
  }
  return true
}

// O(n)
function oneAway(s1: string, s2: string): boolean {
  if (s2.length - s1.length === 1) {
    return hasAdditionalCharacter(s1, s2)
  } else if (s1.length - s2.length === 1) {
    return hasAdditionalCharacter(s2, s1)
  } else if (s1.length === s2.length) {
    let seenReplacement = false
    for (let i = 0; i < s1.length; i++) {
      if (s1[i] !== s2[i]) {
        if (seenReplacement) {
          return false
        }
        seenReplacement = true
      }
    }
    return seenReplacement
  }
  return false
}
