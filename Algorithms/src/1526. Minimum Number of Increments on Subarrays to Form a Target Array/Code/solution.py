from typing import List

class Solution:
    def minNumberOperations(self, target: List[int]) -> int:
        if not target:
            return 0
        ans = target[0]  # operations needed to build target[0] from 0
        for i in range(1, len(target)):
            if target[i] > target[i - 1]:
                ans += target[i] - target[i - 1]  # only positive increases contribute
        return ans
