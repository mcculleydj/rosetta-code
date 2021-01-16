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

# maintain an index pointer for each ordered list initialized to the start
# set the current minimum distance to the absolute difference of the first elements
# increment the index pointer that results in the smaller absolute distance
# note, this may be worse than the current minimum
# keep track of the minimum distance encountered across all iterations
# eventually, one pointer will reach the end of the list
# advance the other pointer so long as current distance is decreasing
# once distance starts increasing it will never get smaller
# since one list has been completely exhausted and the lists are ordered
# return the minimum distance encountered in O(m + n)

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


# O(m + n) find the element in arr1 and the element in arr2
# that are closest to each other and return the distance
def closest_pair(arr1, arr2):
    if len(arr1) == 1 and len(arr2) == 1:
        return abs(arr1[0] - arr2[0])

    # m || n == 1 case could be handled in O(log n) using binary search
    if len(arr1) == 1:
        distance = None
        current_distance = None
        i = 0
        while i < len(arr2) and (current_distance is None or distance < current_distance):
            current_distance = distance
            distance = abs(arr1[0] - arr2[i])
            i += 1
        return min(distance, current_distance)
    elif len(arr2) == 1:
        distance = None
        current_distance = None
        i = 0
        while i < len(arr1) and (current_distance is None or distance < current_distance):
            current_distance = distance
            distance = abs(arr2[0] - arr1[i])
            i += 1
        return min(distance, current_distance)

    i = 0
    j = 0
    min_distance = None
    distance = 0

    while i < len(arr1) - 1 and j < len(arr2) - 1:
        distance = abs(arr1[i] - arr2[j])
        advance_i_distance = abs(arr1[i+1] - arr2[j])
        advance_j_distance = abs(arr1[i] - arr2[j+1])

        if min_distance is None or distance < min_distance:
            min_distance = distance

        if advance_i_distance < advance_j_distance:
            i += 1
        else:
            j += 1

    current_distance = distance

    if i == len(arr1) - 1:
        while j < len(arr2) and distance <= current_distance:
            j += 1
            current_distance = distance
            distance = abs(arr1[i] - arr2[j])
    elif i != len(arr1) - 1:
        while i < len(arr1) and distance <= current_distance:
            i += 1
            current_distance = distance
            distance = abs(arr1[i] - arr2[j])

    return min(min_distance, current_distance)


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

    return closest_pair(idx1, idx2)
