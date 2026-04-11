class Solution:
    def minimumDistance(self, nums: List[int]) -> int:
        positions = {}
        
        # Store all indices for each value
        for i, num in enumerate(nums):
            if num not in positions:
                positions[num] = []
            positions[num].append(i)
        
        ans = float('inf')
        
        # Check every value's index list
        for idx in positions.values():
            if len(idx) < 3:
                continue
            
            # Check every consecutive group of 3 indices
            for i in range(len(idx) - 2):
                distance = 2 * (idx[i + 2] - idx[i])
                ans = min(ans, distance)
        
        return -1 if ans == float('inf') else ans