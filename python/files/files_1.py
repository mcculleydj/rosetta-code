# Approach:
# read the file and create a dictionary
# assuming we can hold this dictionary in memory
# where each word maps to an ordered list of indices
# such that an entry "foo": [1, 123, 1928] denotes that
# the word "foo" is the 1st, 123rd, and 1928th word in the file

# then we can restate the problem as
# given two lists of integers U and V
# determine the minimum absolute difference
# between any element of U and any element of V
# | u_i - v_j | for all u_i in U and v_j in V

# let S be the shorter list with m elements
# let L be the longer list with n elements
# for any s_i in S we can use binary search
# to determine its proper location in L in O(log n)
# if s_i would occur between l_j, and l_k in L
# then min(s_i - l_j, l_k - s_i) is the minimum
# distance for s_i, the minimum of all minimum distances
# for all s_i in S is the minimum distance between these
# two words in the file and can be determined in O(m log n)

# note, at least one of l_j or l_k must exist for all s_i
# however, if s_i is less than all elements of L
# or greater than all elements of L then one will not exist
# in this case only consider l_k - s_i or l_j - s_i
# for the nearest value in L, either the first or last element

from collections import defaultdict


def read_file():
    word_to_indices = defaultdict(list)
    i = 1

    with open('file_1.txt') as f:
        # lazy for a multi-line file
        # in the event of a single-line file Python can still handle
        # files that exceed available memory by using swap
        for line in f:
            words = line.split()

            for j, word in enumerate(words):
                word_to_indices[word].append(i + j)

            i += len(words)

    return word_to_indices


def binary_search(value, arr):
    if value < arr[0]:
        return arr[0] - value
    if value > arr[len(arr) - 1]:
        return value - arr[len(arr) - 1]
    if len(arr) == 2:
        return min(value - arr[0], arr[1] - value)

    i = 0
    j = len(arr) - 1

    while j - i > 1:
        span = j - i
        middle = arr[i + span // 2]

        if value < middle:
            # move upper bound to middle
            j = i + span // 2
        else:
            # move lower bound to middle
            i += span // 2

    return min(value - arr[i], arr[j] - value)


def min_word_distance(w1, w2):
    word_to_indices = read_file()

    idx1 = word_to_indices.get(w1)
    idx2 = word_to_indices.get(w2)

    if not idx1 and not idx2:
        raise Exception(f'neither {w1} or {w2} appear in the file')
    if not idx1:
        raise Exception(f'{w1} does not appear in the file')
    if not idx2:
        raise Exception(f'{w2} does not appear in the file')

    if len(idx1) > len(idx2):
        outer, inner = idx2, idx1
    else:
        outer, inner = idx1, idx2

    min_distance = None

    for i in outer:
        min_i = binary_search(i, inner)

        if min_distance is None or min_i < min_distance:
            min_distance = min_i

    return min_distance
