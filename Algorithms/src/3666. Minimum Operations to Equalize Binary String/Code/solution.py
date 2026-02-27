class Solution:
    def minOperations(self, s: str, k: int) -> int:
        n = len(s)
        zero = s.count('0')

        if zero == 0:
            return 0

        if n == k:
            if zero == n:
                return 1
            if zero == 0:
                return 0
            return -1

        one = n - zero
        base = n - k

        ans = float('inf')

        # Odd operations
        if (k % 2) == (zero % 2):
            m = max(
                (zero + k - 1) // k,
                (one + base - 1) // base
            )

            if m % 2 == 0:
                m += 1

            ans = min(ans, m)

        # Even operations
        if zero % 2 == 0:
            m = max(
                (zero + k - 1) // k,
                (zero + base - 1) // base
            )

            if m % 2 == 1:
                m += 1

            ans = min(ans, m)

        return -1 if ans == float('inf') else ans