class Solution:
    def minCost(self, colors: str, neededTime: List[int]) -> int:
        ans = 0                # total minimal removal time
        block_sum = 0          # sum of times in the current same-color block
        block_max = 0          # max time in the current block
        n = len(colors)
        
        for i in range(n):
            # if color changes, we finalize previous block
            if i > 0 and colors[i] != colors[i-1]:
                ans += block_sum - block_max
                block_sum = 0
                block_max = 0
            block_sum += neededTime[i]
            block_max = max(block_max, neededTime[i])
        
        # add remaining block
        ans += block_sum - block_max
        return ans
