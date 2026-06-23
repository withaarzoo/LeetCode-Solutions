class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1000000007
        m = r - l + 1

        # Length 1: every rank is valid
        dp = [1] * m

        for _ in range(2, n + 1):
            # Reverse so one prefix-sum pass handles
            # the alternating transition automatically
            dp.reverse()

            pref = 0

            for i in range(m):
                old = dp[i]

                # New state gets all previous contributions
                dp[i] = pref

                pref = (pref + old) % MOD

        # Sum all ending positions
        ans = sum(dp) % MOD

        # Count both zigzag directions
        return (ans * 2) % MOD