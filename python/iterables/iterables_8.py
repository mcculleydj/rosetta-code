# O(mn) where m and n are the dimensions of the matrix
def zero_matrix(m):
    zero_rows = set()
    zero_columns = set()

    # O(mn)
    for i in range(len(m)):
        for j in range(len(m[0])):
            if m[i][j] == 0:
                zero_rows.add(i)
                zero_columns.add(j)
    # O(mn)
    for i in range(len(m)):
        for j in range(len(m[0])):
            # O(1) lookup by using sets
            if i in zero_rows or j in zero_columns:
                m[i][j] = 0
