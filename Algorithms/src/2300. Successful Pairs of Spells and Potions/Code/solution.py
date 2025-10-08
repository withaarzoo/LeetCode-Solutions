from bisect import bisect_left
from typing import List

class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()                # sort potions for binary search
        m = len(potions)
        ans = []
        for s in spells:
            # smallest potion value needed (ceil division)
            need = (success + s - 1) // s
            # find first index where potion >= need
            pos = bisect_left(potions, need)
            ans.append(m - pos)
        return ans
