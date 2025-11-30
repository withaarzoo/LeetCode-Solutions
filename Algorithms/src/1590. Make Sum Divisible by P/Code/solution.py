from typing import List

class Solution:
    def minSubarray(self, nums: List[int], p: int) -> int:
        total = 0
        for x in nums:
            total = (total + x) % p  # keep modulo only
        
        need = total
        if need == 0:
            return 0  # already divisible
        
        n = len(nums)
        last_index = {0: -1}  # remainder -> latest index
        ans = n
        prefix = 0
        
        for i, x in enumerate(nums):
            prefix = (prefix + x) % p
            target = (prefix - need) % p  # Python mod handles negative
            
            if target in last_index:
                ans = min(ans, i - last_index[target])
            
            # update latest index of this remainder
            last_index[prefix] = i
        
        return -1 if ans == n else ans
