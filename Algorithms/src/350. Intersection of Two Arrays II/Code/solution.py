from collections import Counter

class Solution:
    def intersect(self, nums1, nums2):
        # Step 1: Create a Counter object from nums1 to count occurrences of each number
        countMap = Counter(nums1)
        
        # Step 2: Initialize an empty list to store the intersection result
        result = []

        # Step 3: Iterate through each number in nums2
        for num in nums2:
            # Step 4: Check if the number exists in countMap and has a count greater than 0
            if countMap[num] > 0:
                # Step 5: If so, append the number to the result list
                result.append(num)
                # Step 6: Decrease the count of the number in countMap
                countMap[num] -= 1

        # Step 7: Return the intersection result list
        return result
