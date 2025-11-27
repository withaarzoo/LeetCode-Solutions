from typing import List

class Solution:
    def maxSubarraySum(self, nums: List[int], k: int) -> int:
        n = len(nums)
        INF = 4 * 10**18  # large number
        
        # minPref[r] = smallest prefix sum seen so far where index % k == r
        minPref = [INF] * k
        
        prefix = 0
        ans = -INF
        
        # prefix index 0 has sum 0 and remainder 0
        minPref[0] = 0
        
        for i, val in enumerate(nums):
            prefix += val
            rem = (i + 1) % k  # prefix index is i+1
            
            if minPref[rem] != INF:
                ans = max(ans, prefix - minPref[rem])
            
            if prefix < minPref[rem]:
                minPref[rem] = prefix
        
        return ans
