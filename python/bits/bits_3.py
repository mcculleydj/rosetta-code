# Approach:
# represent a number between 0 and 1 as the following sum:
# (1 || 0) * 1/2^0 + (1 || 0) * 1/2^1...
# 0.625 base 10 => 0.101 base 2
# assume we start counting characters to reach the max of 32
# after the decimal point

def binary_to_string(n):
    # based on the prompt
    # assume n is between 0 and 1 non-inclusive
    bits = []
    while n > 0:
        if len(bits) > 31:
            raise Exception('loss of precision')

        k = n * 2
        if k >= 1:
            bits.append('1')
            n = k - 1
        else:
            bits.append('0')
            n = k

    return ''.join(bits)
