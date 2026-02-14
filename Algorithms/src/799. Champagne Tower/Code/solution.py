class Solution:
    def champagneTower(self, poured: int, query_row: int, query_glass: int) -> float:
        
        # Create DP table
        dp = [[0.0] * 101 for _ in range(101)]
        dp[0][0] = poured
        
        # Simulate champagne flow
        for r in range(query_row + 1):
            for c in range(r + 1):
                if dp[r][c] > 1.0:
                    overflow = (dp[r][c] - 1.0) / 2.0
                    dp[r + 1][c] += overflow
                    dp[r + 1][c + 1] += overflow
                    dp[r][c] = 1.0
        
        return min(1.0, dp[query_row][query_glass])
