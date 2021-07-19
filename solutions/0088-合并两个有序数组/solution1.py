from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        p1, p2, ptr = m - 1, n - 1, m + n - 1
        while ptr >= 0:
            if p1 == -1:
                nums1[ptr] = nums2[p2]
                p2 -= 1
            elif p2 == -1:
                nums1[ptr] = nums1[p1]
                p1 -= 1
            elif nums1[p1] > nums2[p2]:
                nums1[ptr] = nums1[p1]
                p1 -= 1
            else:
                nums1[ptr] = nums2[p2]
                p2 -= 1
            ptr -= 1
