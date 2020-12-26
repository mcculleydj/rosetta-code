# O(n)
# could rely on built in replace method, but that seems to be what the problem is asking us to implement
def url_encode(s):
    # strings are immutable so building a list and using ''.join()
    # is a better approach than building a new string with +=

    # O(n) for the list comprehension and O(n) for join()
    return ''.join(['%20' if c == ' ' else c for c in s])
