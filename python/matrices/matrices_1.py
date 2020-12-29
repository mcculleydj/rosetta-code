# O(n^2) for an n x n matrix
def rotate_matrix(m):
    # given a cell (i, j) the destination cell is (j, n - i - 1) for an n x n matrix

    # assume m is represented by a 2D list of integers
    # nested for loop is O(n^2)
    # in this loop "i" represents the layer
    # in an n x n matrix there are floor(n / 2) layers to rotate
    for i in range(len(m) // 2):
        for j in range(i, len(m) - i - 1):
            # results in 4 recursive calls
            rotate_cell(m, m[i][j], (j, len(m) - i - 1), (i, j))


# rotate_cell is a recursive helper that sets the value of
# of the destination cell to the supplied value
# rotate_cell will be called 4 times before destination == start
def rotate_cell(m, value, destination, start):
    previous_value = m[destination[0]][destination[1]]
    m[destination[0]][destination[1]] = value

    # base case
    if destination == start:
        return

    rotate_cell(m, previous_value,
                (destination[1], len(m) - destination[0] - 1), start)
