func stoneGameIII(stoneValue []int) string {

	n := len(stoneValue)

	// dp[i] stores the maximum score difference from index i
	dp := make([]int, n+1)

	// Process from right to left
	for i := n - 1; i >= 0; i-- {

		// Very small initial value
		dp[i] = -(1 << 30)

		sum := 0

		// Try taking 1, 2 and 3 stones
		for j := i; j < n && j < i+3; j++ {

			// Add current stone value
			sum += stoneValue[j]

			// Keep the maximum score difference
			if sum-dp[j+1] > dp[i] {
				dp[i] = sum - dp[j+1]
			}
		}
	}

	// Decide the winner
	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}