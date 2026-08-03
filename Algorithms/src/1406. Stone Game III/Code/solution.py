class Solution:
    def stoneGameIII(self, stoneValue: List[int]) -> str:

        n = len(stoneValue)

        # dp[i] stores the maximum score difference from index i
        dp = [0] * (n + 1)

        # Fill DP from back to front
        for i in range(n - 1, -1, -1):

            dp[i] = float("-inf")
            total = 0

            # Try taking 1, 2 and 3 stones
            for j in range(i, min(n, i + 3)):

                # Add current stone value
                total += stoneValue[j]

                # Update the best score difference
                dp[i] = max(dp[i], total - dp[j + 1])

        # Decide the winner
        if dp[0] > 0:
            return "Alice"
        elif dp[0] < 0:
            return "Bob"
        else:
            return "Tie"