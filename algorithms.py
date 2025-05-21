def bubble_sort(data: list) -> list:
    """Sorts a list in ascending order using the Bubble Sort algorithm.

    Args:
        data: The list to be sorted.

    Returns:
        A new list containing the elements of `data` sorted in ascending order.
    """
    # Create a copy of the list to avoid modifying the original
    new_list = data[:]
    n = len(new_list)
    for i in range(n):
        # Last i elements are already in place
        for j in range(0, n-i-1):
            # Swap if the element found is greater than the next element
            if new_list[j] > new_list[j+1]:
                new_list[j], new_list[j+1] = new_list[j+1], new_list[j]
    return new_list

def quick_sort(data: list) -> list:
    """Sorts a list in ascending order using the Quick Sort algorithm.

    Args:
        data: The list to be sorted.

    Returns:
        A new list containing the elements of `data` sorted in ascending order.
    """
    if len(data) <= 1:
        return data[:]  # Return a copy for consistency, even for base cases
    else:
        # Choose the last element as the pivot
        pivot = data[-1]
        # Partition the list (excluding the pivot itself)
        lesser = [x for x in data[:-1] if x <= pivot]
        greater = [x for x in data[:-1] if x > pivot]
        # Recursively sort the sub-arrays and combine them
        return quick_sort(lesser) + [pivot] + quick_sort(greater)
