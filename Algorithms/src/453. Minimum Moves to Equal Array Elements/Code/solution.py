class Solution:
    def minMoves(self, nums: list[int]) -> int:
        mn = min(nums)
        return sum(x - mn for x in nums)