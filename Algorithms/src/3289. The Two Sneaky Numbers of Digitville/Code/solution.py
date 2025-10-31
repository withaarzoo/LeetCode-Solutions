from typing import List

class Solution:
    def getSneakyNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums) - 2                   # original range size
        seen = [False] * n                  # boolean marks for 0..n-1
        res = []
        for x in nums:
            if seen[x]:
                res.append(x)               # found a duplicate
                if len(res) == 2:
                    break                   # early exit
            else:
                seen[x] = True              # mark seen
        return res
