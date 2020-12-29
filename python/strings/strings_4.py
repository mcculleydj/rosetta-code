# O(n)
def palindrome_permutation(s):
    character_count = {}

    # O(n)
    for c in s:
        if character_count.get(c):
            character_count[c] += 1
        else:
            character_count[c] = 1

    # a string is a palindrome if all character counts are even
    # if any are odd counts, there can be only one serving as the middle character

    has_odd = False

    # number of unique characters is always <= number of characters still O(n)
    for v in character_count.values():
        if v % 2 == 1:
            if has_odd:
                # more than one odd
                return False
            has_odd = True

    return True
