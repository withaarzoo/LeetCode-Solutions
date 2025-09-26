from typing import List

class Solution:
    def triangleNumber(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 3:
            return 0
        nums.sort()                                # sort ascending
        count = 0
        # k is index of the largest side (c)
        for k in range(n - 1, 1, -1):
            i, j = 0, k - 1                        # two pointers for a and b
            while i < j:
                if nums[i] + nums[j] > nums[k]:
                    count += j - i                 # every index between i and j-1 forms a pair with j
                    j -= 1                         # move to smaller b to count other combos
                else:
                    i += 1                         # need larger a, move left pointer right
        return count
