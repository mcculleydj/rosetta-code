# Approach:
# Shift M such that its right-most bit is at the ith position
# Mask N such that all bits between i and i + most_significant_bit(M) is 0
# Use a bitwise OR to produce the desired number


def get_most_significant_bit_32(n):
    if n == 0:
        return 0

    msb = 0

    for i in range(32):
        b = n & 1
        n >>= 1

        if b == 1:
            msb = i

    return msb


def insert(m, n, i):
    msb = get_most_significant_bit_32(m)
    m <<= i

    # preserve N to the left of the range to replace (1s)
    left = -1 << i + msb
    # preserve N to the right of i (1s)
    right = (1 << i) - 1
    # O from i to i + msb and 1 outside of this range
    mask = left | right
    # mask N
    n &= mask

    return m | n
