# Python3 (bisect on a growing list of dry days)
from typing import List
import bisect

class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        n = len(rains)
        ans = [1] * n
        last = {}  # lake -> last day index
        dry = []   # sorted list of dry day indices (increasing because we append)

        for i, lake in enumerate(rains):
            if lake > 0:
                ans[i] = -1
                if lake in last:
                    prev = last[lake]
                    # find the first dry day strictly greater than prev
                    idx = bisect.bisect_right(dry, prev)
                    if idx == len(dry):
                        return []
                    dry_day = dry[idx]
                    ans[dry_day] = lake
                    dry.pop(idx)  # remove used dry day
                last[lake] = i
            else:
                dry.append(i)  # append keeps it sorted (i increases each loop)
        return ans
