class Solution:
    def minimumDistance(self, nums: List[int]) -> int:
        pos = {}
        
        # Store all indices for each value
        for i, num in enumerate(nums):
            if num not in pos:
                pos[num] = []
            pos[num].append(i)
        
        ans = float('inf')
        
        # Process each value's indices
        for indices in pos.values():
            if len(indices) < 3:
                continue
            
            # Check every consecutive group of 3 indices
            for i in range(len(indices) - 2):
                distance = 2 * (indices[i + 2] - indices[i])
                ans = min(ans, distance)
        
        return -1 if ans == float('inf') else ans