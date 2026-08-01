func predictTheWinner(nums []int) bool {
	n := len(nums)

	// dp[i][j] stores the maximum score difference
	// current player can achieve for subarray [i...j].
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case:
	// Only one number remains.
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}

	// Fill DP for larger subarrays.
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			// Take the left number.
			takeLeft := nums[i] - dp[i+1][j]

			// Take the right number.
			takeRight := nums[j] - dp[i][j-1]

			// Store the better choice.
			if takeLeft > takeRight {
				dp[i][j] = takeLeft
			} else {
				dp[i][j] = takeRight
			}
		}
	}

	// Player 1 wins or ties.
	return dp[0][n-1] >= 0
}