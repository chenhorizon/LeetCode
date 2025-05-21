import unittest
from algorithms import bubble_sort, quick_sort

class TestBubbleSort(unittest.TestCase):
    def test_empty_list(self):
        self.assertEqual(bubble_sort([]), [])

    def test_sorted_list(self):
        self.assertEqual(bubble_sort([1, 2, 3, 4, 5]), [1, 2, 3, 4, 5])

    def test_reverse_sorted_list(self):
        self.assertEqual(bubble_sort([5, 4, 3, 2, 1]), [1, 2, 3, 4, 5])

    def test_duplicate_elements(self):
        self.assertEqual(bubble_sort([3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]), [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9])

    def test_negative_numbers(self):
        self.assertEqual(bubble_sort([-3, -1, -4, -1, -5]), [-5, -4, -3, -1, -1])

    def test_single_element_list(self):
        self.assertEqual(bubble_sort([42]), [42])

    def test_mixed_numbers(self):
        self.assertEqual(bubble_sort([3, -1, 0, 5, -2]), [-2, -1, 0, 3, 5])

class TestQuickSort(unittest.TestCase):
    def test_empty_list(self):
        self.assertEqual(quick_sort([]), [])

    def test_sorted_list(self):
        self.assertEqual(quick_sort([1, 2, 3, 4, 5]), [1, 2, 3, 4, 5])

    def test_reverse_sorted_list(self):
        self.assertEqual(quick_sort([5, 4, 3, 2, 1]), [1, 2, 3, 4, 5])

    def test_duplicate_elements(self):
        self.assertEqual(quick_sort([3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]), [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9])

    def test_negative_numbers(self):
        self.assertEqual(quick_sort([-3, -1, -4, -1, -5]), [-5, -4, -3, -1, -1])

    def test_single_element_list(self):
        self.assertEqual(quick_sort([42]), [42])

    def test_mixed_numbers(self):
        self.assertEqual(quick_sort([3, -1, 0, 5, -2]), [-2, -1, 0, 3, 5])

if __name__ == '__main__':
    unittest.main()
