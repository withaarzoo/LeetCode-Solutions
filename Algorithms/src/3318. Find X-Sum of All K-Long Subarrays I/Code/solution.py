from collections import defaultdict
from typing import List

class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        n = len(nums)
        freq = defaultdict(int)

        # initial window
        for i in range(k):
            freq[nums[i]] += 1

        def compute_x_sum(freq, x):
            # items: (value, freq)
            items = [(v, f) for v, f in freq.items()]
            # sort by freq desc, value desc
            items.sort(key=lambda t: (-t[1], -t[0]))
            total = 0
            for i in range(min(x, len(items))):
                v, f = items[i]
                total += v * f
            return total

        ans = [compute_x_sum(freq, x)]

        # slide the window
        for i in range(k, n):
            add = nums[i]
            rem = nums[i - k]
            freq[add] += 1
            freq[rem] -= 1
            if freq[rem] == 0:
                del freq[rem]
            ans.append(compute_x_sum(freq, x))

        return ans
