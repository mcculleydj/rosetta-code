// O(n^2)
function rotateMatrix(m: number[][]) {
  // nested loop is O(n^2)
  for (let i = 0; i < Math.floor(m.length / 2); i++) {
    for (let j = i; j < m.length - i - 1; j++) {
      rotateCell(m, m[i][j], [j, m.length - i - 1], [i, j])
    }
  }
}

// O(1) for each rotation with 4 recursive calls
function rotateCell(
  m: number[][],
  value: number,
  destination: number[],
  start: number[],
) {
  const previousValue = m[destination[0]][destination[1]]
  m[destination[0]][destination[1]] = value

  if (destination[0] === start[0] && destination[1] === start[1]) {
    return
  }

  rotateCell(
    m,
    previousValue,
    [destination[1], m.length - destination[0] - 1],
    start,
  )
}
