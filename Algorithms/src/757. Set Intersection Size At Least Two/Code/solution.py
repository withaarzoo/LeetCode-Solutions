from typing import List

class Solution:
    def intersectionSizeTwo(self, intervals: List[List[int]]) -> int:
        # sort by end ascending, for same end sort start descending
        intervals.sort(key=lambda x: (x[1], -x[0]))
        
        a = -10**18  # smaller of the last two chosen
        b = -10**18  # larger of the last two chosen
        ans = 0
        
        for l, r in intervals:
            if l > b:
                # none of a,b in [l,r], pick r-1 and r
                ans += 2
                a = r - 1
                b = r
            elif l > a:
                # only b is in [l,r], pick r
                ans += 1
                a = b
                b = r
            else:
                # both are already inside
                pass
        return ans
