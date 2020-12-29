# O(n)
def is_rotation(s1, s2):
    if len(s1) != len(s2):
        return False

    # let AB represent s1
    # where A is the first part of the string and B is the second
    # s2 is a rotation <=> s2 == BA (necessary and sufficient)
    # let s2 concatenated be represented by BABA
    # if s1 can be represented by AB then it must be a substring of s2 + s2

    # O(n)
    return s1 in s2 + s2
