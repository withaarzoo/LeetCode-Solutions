class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        # Use a set for O(1) average membership checks
        s = set(nums)
        # While original exists in the set, keep doubling it
        while original in s:
            original *= 2
        return original
