from typing import List

class Solution:
    def lexicographicallySmallestArray(self, nums: List[int], limit: int) -> List[int]:
        n = len(nums)  # Store the size of the array.
        
        # Store each value together with its original index.
        elements = [(nums[i], i) for i in range(n)]
        
        # Sort by value so connected values become consecutive.
        elements.sort()
        
        # Store the lexicographically smallest result.
        answer = [0] * n
        
        start = 0  # First element of the current connected group.
        
        while start < n:
            end = start  # Expand the current group.
            
            # Consecutive values belong to the same group when their
            # difference is at most limit.
            while (
                end + 1 < n
                and elements[end + 1][0] - elements[end][0] <= limit
            ):
                end += 1
            
            # Collect all original indices belonging to this group.
            indices = []
            
            for i in range(start, end + 1):
                indices.append(elements[i][1])
            
            # Sort positions so smaller values go to earlier positions.
            indices.sort()
            
            # The values are already sorted because elements was sorted.
            for i, index in enumerate(indices):
                answer[index] = elements[start + i][0]
            
            # Move to the next connected group.
            start = end + 1
        
        return answer  # Return the final lexicographically smallest array.