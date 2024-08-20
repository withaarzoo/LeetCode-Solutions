func stoneGameII(piles []int) int {
    // Get the number of piles
    n := len(piles)

    // Initialize the DP table where dp[i][M] will store the maximum number of stones the current player can get 
    // starting from the ith pile with the current M value
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // Initialize the suffixSum array where suffixSum[i] will store the total number of stones from pile i to the end
    suffixSum := make([]int, n+1)

    // Calculate the suffix sums in reverse order
    // suffixSum[i] represents the total number of stones from the i-th pile to the last pile
    for i := n - 1; i >= 0; i-- {
        suffixSum[i] = suffixSum[i+1] + piles[i]
    }

    // Fill the DP table by processing each pile from the last one back to the first
    // dp[i][M] represents the maximum stones the current player can collect starting from pile i with a current M value
    for i := n - 1; i >= 0; i-- {
        // Iterate through possible values of M, where M starts from 1 and can go up to n
        for M := 1; M <= n; M++ {
            // Iterate through the possible choices of X, where the current player can take 1 to 2*M piles
            // Ensure that the player does not take more piles than available (i + X <= n)
            for X := 1; X <= 2*M && i+X <= n; X++ {
                // The current player wants to maximize their score, hence subtract the opponent's best result from the total
                dp[i][M] = max(dp[i][M], suffixSum[i]-dp[i+X][max(M, X)])
            }
        }
    }

    // The answer is the maximum number of stones the first player can get starting from the first pile with M=1
    return dp[0][1]
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
