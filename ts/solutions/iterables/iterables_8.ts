// O(mn)
function zeroMatrix(m: number[][]) {
  const zeroRows: Set<number> = new Set()
  const zeroColumns: Set<number> = new Set()

  for (let i = 0; i < m.length; i++) {
    for (let j = 0; j < m[0].length; j++) {
      if (m[i][j] === 0) {
        zeroRows.add(i)
        zeroColumns.add(j)
      }
    }
  }

  for (let i = 0; i < m.length; i++) {
    for (let j = 0; j < m[0].length; j++) {
      if (zeroRows.has(i) || zeroColumns.has(j)) {
        m[i][j] = 0
      }
    }
  }
}
