class Solution:
    def repeatedNTimes(self, nums: List[int]) -> int:
        seen = set()

        for x in nums:
            # If number already seen, return it
            if x in seen:
                return x
            # Otherwise, add to set
            seen.add(x)
