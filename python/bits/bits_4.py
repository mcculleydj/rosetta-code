# Approach:
# Pass over the binary repr storing the index of every 0
# Also store the length of the repr (MSB)
# Handle no zeroes and single zero edge cases by checking length of index array
# 0 to index of second 0 is the first range of 1s that can be produced by flipping a bit
# Set that as the longest seen thus far and then iterate over the array
# Look ahead to calculate all ranges and update longest seen accordingly
# Break the loop when the current zero is the next to last zero in the repr

def flip_bit(n):
    # current index, then used as the length of the binary repr
    i = 0
    # all indices at which a 0 appears
    zero_indices = []

    # O(number of bits) -- will be language / arch dependent
    while n > 0:
        if n & 1 == 0:
            zero_indices.append(i)
        n >>= 1
        i += 1

    # handle two edge cases
    if len(zero_indices) == 0:
        # all 1s => flip the first or last to 0 for i - 1
        return i - 1
    elif len(zero_indices) == 1:
        # single 0 => flip the 0 for i
        return i

    # start with the index of the second zero
    # flipping the first 0 results in a series of 1s
    # from 0 to this index
    longest = zero_indices[1]

    for j, idx in enumerate(zero_indices):
        # if this is the last 0 in the repr we're done
        if len(zero_indices) - 1 == j:
            break

        # set starting index one beyond this 0
        start = idx + 1

        # look ahead
        if len(zero_indices) - 2 == j:
            # if this is the next to last 0 this will be the final range
            end = i
        else:
            # set the end point to the zero after the zero in the middle of this range (+2)
            end = zero_indices[j + 2]

        # store the longest range seen thus far
        if end - start > longest:
            longest = end - start

    return longest
