from functools import cache

class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        n = len(s)
        masks = [1 << (ord(c) - ord('a')) for c in s]

        @cache
        def dp(i, can_change, mask):
            if i == n:
                return 0
            new_mask = mask | masks[i]
            res = 0

            if new_mask.bit_count() > k:
                res = 1 + dp(i + 1, can_change, masks[i])
            else:
                res = dp(i + 1, can_change, new_mask)

            if can_change:
                for j in range(26):
                    changed_mask = mask | (1 << j)
                    if changed_mask.bit_count() > k:
                        res = max(res, 1 + dp(i + 1, 0, 1 << j))
                    else:
                        res = max(res, dp(i + 1, 0, changed_mask))
            return res

        return dp(0, 1, 0) + 1
