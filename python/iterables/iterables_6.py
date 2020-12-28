# O(n)
def compress_string(s):
    # handle edge case ahead of assignment to s[0]
    if s == '':
        return ''

    # accumulator list to store compressed string
    acc = []

    start_index = 0
    current_index = 0
    current_char = s[0]

    # O(n)
    for c in s:
        if c != current_char:
            # O(1)
            acc.extend([current_char, str(current_index - start_index)])
            start_index = current_index
            current_char = c
        current_index += 1

    acc.extend([current_char, str(current_index - start_index)])

    # O(n)
    return ''.join(acc) if len(acc) < len(s) else s
