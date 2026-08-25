from typing import List

class Solution:
    def missingMultiple(self, nums: List[int], k: int) -> int:
        # Store all values so checking whether a multiple exists is fast.
        present = set(nums)

        # Start with the smallest positive multiple of k.
        multiple = k

        # Continue while the current multiple already exists.
        while multiple in present:
            # Move to the next positive multiple of k.
            multiple += k

        # Return the first missing multiple.
        return multiple