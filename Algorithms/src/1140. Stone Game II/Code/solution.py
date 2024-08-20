class Solution:
    def stoneGameII(self, piles: List[int]) -> int:
        n = len(piles)  # Get the total number of piles
        # Initialize a DP table where dp[i][M] represents the maximum number of stones the current player can get
        # starting from pile i with a maximum of M piles to take.
        dp = [[0] * (n + 1) for _ in range(n + 1)]
        
        # Initialize a suffixSum array where suffixSum[i] represents the total number of stones from pile i to the end.
        suffixSum = [0] * (n + 1)
        
        # Compute the suffix sums from the last pile to the first.
        # This will help in calculating the remaining stones efficiently.
        for i in range(n - 1, -1, -1):
            suffixSum[i] = suffixSum[i + 1] + piles[i]
        
        # Fill the DP table by calculating the maximum stones the player can get from each state.
        for i in range(n - 1, -1, -1):  # Iterate from the last pile to the first.
            for M in range(1, n + 1):  # Iterate over all possible values of M.
                for X in range(1, min(2 * M, n - i) + 1):  # Try taking X piles where 1 <= X <= min(2*M, remaining piles).
                    # Update dp[i][M] with the maximum stones the player can get.
                    dp[i][M] = max(dp[i][M], suffixSum[i] - dp[i + X][max(M, X)])
        
        # Return the result from the starting point with M = 1, which is the maximum number of stones the first player can get.
        return dp[0][1]
