import { readFileSync } from 'fs'

function readFile(path: string): { [key: string]: number[] } {
  const file = readFileSync(path, 'utf-8')
  const wordToIndices: { [key: string]: number[] } = {}

  file.split(/\s+/).forEach((w, i) => {
    if (wordToIndices[w]) {
      wordToIndices[w].push(i + 1)
    } else {
      wordToIndices[w] = [i + 1]
    }
  })

  return wordToIndices
}

// O(m + n) find the element in idx1 and the element in idx2
// that are closest to each other and return the distance
function closestPair(idx1: number[], idx2: number[]): number {
  if (idx1.length === 1 && idx2.length === 1) {
    return Math.abs(idx1[0] - idx2[0])
  }

  if (idx1.length === 1) {
    let distance = 0
    let currentDistance = 0
    let i = 0

    while (
      i < idx2.length &&
      (currentDistance === 0 || distance < currentDistance)
    ) {
      currentDistance = distance
      distance = Math.abs(idx1[0] - idx2[i])
      i++
    }

    return Math.min(distance, currentDistance)
  } else if (idx2.length === 1) {
    let distance = 0
    let currentDistance = 0
    let i = 0

    while (
      i < idx1.length &&
      (currentDistance === 0 || distance < currentDistance)
    ) {
      currentDistance = distance
      distance = Math.abs(idx2[0] - idx1[i])
      i++
    }

    return Math.min(distance, currentDistance)
  }

  let i = 0
  let j = 0
  let minDistance = -1
  let distance = 0

  while (i < idx1.length - 1 && j < idx2.length - 1) {
    distance = Math.abs(idx1[i] - idx2[j])
    const iDistance = Math.abs(idx1[i + 1] - idx2[j])
    const jDistance = Math.abs(idx1[i] - idx2[j + 1])

    if (minDistance === -1 || distance < minDistance) {
      minDistance = distance
    }

    if (iDistance < jDistance) {
      i++
    } else {
      j++
    }
  }

  let currentDistance = distance

  if (i === idx1.length - 1) {
    while (j < idx2.length && distance <= currentDistance) {
      j++
      currentDistance = distance
      distance = Math.abs(idx1[i] - idx2[j])
    }
  } else if (i !== idx1.length - 1) {
    while (i < idx1.length && distance <= currentDistance) {
      i++
      currentDistance = distance
      distance = Math.abs(idx1[i] - idx2[j])
    }
  }

  return Math.min(minDistance, currentDistance)
}

function minWordDistance(
  wordsToIndices: { [key: string]: number[] },
  w1: string,
  w2: string,
): number {
  const idx1 = wordsToIndices[w1]
  const idx2 = wordsToIndices[w2]

  if (!idx1 && !idx2) {
    throw `neither ${w1} or ${w2} appear in the file`
  }
  if (!idx1) {
    throw `${w1} does not appear in the file`
  }
  if (!idx2) {
    throw `${w2} does not appear in the file`
  }

  return closestPair(idx1, idx2)
}
