class Solution:
    def backtrack(self, nums, index, currentOR, maxOR, count):
        if currentOR == maxOR:
            count[0] += 1
        
        for i in range(index, len(nums)):
            self.backtrack(nums, i + 1, currentOR | nums[i], maxOR, count)
    
    def countMaxOrSubsets(self, nums: List[int]) -> int:
        maxOR = 0
        
        # Step 1: Compute the maximum OR
        for num in nums:
            maxOR |= num
        
        count = [0]
        # Step 2: Backtrack to count the subsets
        self.backtrack(nums, 0, 0, maxOR, count)
        
        return count[0]
