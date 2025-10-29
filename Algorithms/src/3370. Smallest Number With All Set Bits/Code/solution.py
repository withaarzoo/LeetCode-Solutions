class Solution:
    def smallestNumber(self, n: int) -> int:
        # Find smallest k such that (2^k - 1) >= n
        k = 1
        while True:
            val = (1 << k) - 1  # 2**k - 1
            if val >= n:
                return val
            k += 1
