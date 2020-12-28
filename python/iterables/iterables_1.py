# with access to O(1) lookup data structures O(n)
def is_unique(s):
    char_set = set()
    for c in s:
        if c in char_set:
            return False
        char_set.add(c)
    return True


# without access to any additional data structures O(n^2)
# could sort in O(n log(n)), but the assumption is that a new string
# constitutes an additional data structure
def is_unique_slow(s):
    for i in range(len(s)):
        for j in range(i+1, len(s)):
            if s[i] == s[j]:
                return False
    return True
