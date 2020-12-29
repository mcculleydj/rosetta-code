# see Remove Duplicates for List and Node type definitions


# O(n)
def kth_to_last_element(linked_list, k):
    # maintain a leading and lagging pointer
    # lag will trail the lead by k
    lead_node = linked_list.head
    lag_node = linked_list.head
    lead_index = 0
    lag_index = 0

    while lead_node:
        lead_node = lead_node.next
        lead_index += 1
        if lead_index > k:
            lag_node = lag_node.next
            lag_index += 1

    if lead_index - lag_index < k:
        print('List is not long enough to have a %d(st/nd/rd/th) to last element, returning the first node instead.')

    return lag_node
