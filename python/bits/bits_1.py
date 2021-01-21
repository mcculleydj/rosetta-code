# will not overflow in Python because integers have arbitrary precision
# under the covers this may require arithmetic operations
# but, the spirit of the prompt is to solve this using bitwise operators


def add(x, y):
    while y != 0:
        # produce a bit vector where both bits are 1s
        # and we will need to carry over to the left
        carry = x & y

        # XOR is the result of addition without carrying
        x ^= y

        # shift the carry vector and repeat this process
        y = carry << 1

    return x
