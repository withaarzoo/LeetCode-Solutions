class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        i = 0
        j = 0
        ans = 0

        while i < len(nums1) and j < len(nums2):

            # Ensure i <= j
            if i > j:
                j += 1
                continue

            # Valid pair
            if nums1[i] <= nums2[j]:
                ans = max(ans, j - i)
                j += 1  # Try for larger distance
            else:
                # Invalid pair
                i += 1

        return ans