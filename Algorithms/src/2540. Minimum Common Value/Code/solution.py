class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        
        # Pointer for nums1
        i = 0

        # Pointer for nums2
        j = 0

        # Traverse both arrays together
        while i < len(nums1) and j < len(nums2):

            # If both values are same,
            # return the common value
            if nums1[i] == nums2[j]:
                return nums1[i]

            # Move the pointer with smaller value
            if nums1[i] < nums2[j]:
                i += 1
            else:
                j += 1

        # No common value exists
        return -1