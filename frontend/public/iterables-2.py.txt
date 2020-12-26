# O(n + m)
def check_permutation(s1, s2):
    # need to be of equal length O(1)
    if len(s1) != len(s2):
        return False

    # for O(1) lookup
    char_count = {}

    # O(n) where n is the length of s1
    for c in s1:
        if char_count.get(c):
            char_count[c] += 1
        else:
            char_count[c] = 1

    # O(m) where m is the length of s2
    for c in s2:
        # relies on 0 being falsy
        if not char_count.get(c):
            return False
        char_count[c] -= 1

    return True
