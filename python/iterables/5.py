# O(n) determines if s2 is s1 with an additional character
def has_additional_character(s1, s2):
    # can skip a single additional character in s2 to create s1
    has_skipped = False
    i = 0
    j = 0

    while i < len(s1):
        if s1[i] == s2[j]:
            i += 1
            j += 1
        elif not has_skipped:
            has_skipped = True
            j += 1
        else:
            return False

    return True


# O(n)
def one_away(s1, s2):
    # O(n) add case
    if len(s2) - len(s1) == 1:
        return has_additional_character(s1, s2)

    # O(n) remove case
    elif len(s1) - len(s2) == 1:
        return has_additional_character(s2, s1)

    # O(n) handle replace
    elif len(s1) == len(s2):
        # can have at most one replacement
        seen_replacement = False

        for i in range(len(s1)):
            if s1[i] != s2[i]:
                if seen_replacement:
                    return False
                seen_replacement = True

        # will be False if strings are identical
        return seen_replacement

    # if string lengths differ by more than 1
    return False
