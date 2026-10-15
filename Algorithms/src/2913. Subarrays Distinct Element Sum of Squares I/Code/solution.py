class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        n = len(nums)
        res = 0
        for i in range(n):
            s = set()
            for j in range(i, n):
                s.add(nums[j])
                d = len(s)
                res += d * d
        return res