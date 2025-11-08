class Solution:
    def minimumOneBitOperations(self, n: int) -> int:
        # Inverse Gray code: accumulate XORs of n, n>>1, n>>2, ...
        ans = 0
        while n:
            ans ^= n
            n >>= 1
        return ans
