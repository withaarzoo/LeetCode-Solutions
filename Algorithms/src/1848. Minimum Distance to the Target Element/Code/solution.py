class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        # Store the minimum distance found so far
        answer = float('inf')

        # Traverse through the array
        for i in range(len(nums)):
            # Check if current element is the target
            if nums[i] == target:
                # Update the minimum distance
                answer = min(answer, abs(i - start))

        return answer