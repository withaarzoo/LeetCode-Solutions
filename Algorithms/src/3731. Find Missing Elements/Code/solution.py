class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        # Store every number for fast lookup
        seen = set(nums)

        # Find the smallest and largest values
        mn = min(nums)
        mx = max(nums)

        # Store the missing numbers
        ans = []

        # Check every number in the range
        for x in range(mn, mx + 1):
            # If the number is missing, add it
            if x not in seen:
                ans.append(x)

        return ans