from typing import List

class Solution:
    def maxSumDivThree(self, nums: List[int]) -> int:
        total = 0
        INF = 10**9
        r1_min1 = INF  # smallest remainder-1
        r1_min2 = INF  # second smallest remainder-1
        r2_min1 = INF  # smallest remainder-2
        r2_min2 = INF  # second smallest remainder-2

        for x in nums:
            total += x
            r = x % 3
            if r == 1:
                if x < r1_min1:
                    r1_min2 = r1_min1
                    r1_min1 = x
                elif x < r1_min2:
                    r1_min2 = x
            elif r == 2:
                if x < r2_min1:
                    r2_min2 = r2_min1
                    r2_min1 = x
                elif x < r2_min2:
                    r2_min2 = x

        mod = total % 3
        if mod == 0:
            return total

        remove_cost = 10**18

        if mod == 1:
            if r1_min1 != INF:
                remove_cost = min(remove_cost, r1_min1)
            if r2_min2 != INF:
                remove_cost = min(remove_cost, r2_min1 + r2_min2)
        else:  # mod == 2
            if r2_min1 != INF:
                remove_cost = min(remove_cost, r2_min1)
            if r1_min2 != INF:
                remove_cost = min(remove_cost, r1_min1 + r1_min2)

        if remove_cost >= 10**18:
            return 0
        return total - remove_cost
