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

function binarySearch(n: number, idx: number[]): number {
  if (idx.length === 0) {
    throw 'index array is empty'
  }

  if (n < idx[0]) return idx[0] - n
  if (n > idx[idx.length - 1]) return n - idx[idx.length - 1]
  if (idx.length === 2) return Math.min(n - idx[0], idx[1] - n)

  let i = 0
  let j = idx.length - 1
  let span = j - i

  while (span > 1) {
    const middle = idx[Math.floor(i + span / 2)]

    if (n < middle) {
      j = Math.floor(i + span / 2)
    } else {
      i += Math.floor(span / 2)
    }

    span = j - i
  }

  return Math.min(n - idx[i], idx[j] - n)
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

  let outer: number[]
  let inner: number[]

  if (idx1.length > idx2.length) {
    outer = idx2
    inner = idx1
  } else {
    outer = idx1
    inner = idx2
  }

  let minDistance: number

  outer.forEach(n => {
    const min = binarySearch(n, inner)
    if (!minDistance || min < minDistance) {
      minDistance = min
    }
  })

  return minDistance
}
