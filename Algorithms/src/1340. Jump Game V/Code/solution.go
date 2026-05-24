func maxJumps(arr []int, d int) int {

	n := len(arr)

	// dp[i] stores maximum jumps starting from i
	dp := make([]int, n)

	// Initialize with -1 meaning unvisited
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	// DFS function
	var dfs func(int) int

	dfs = func(i int) int {

		// Return stored answer
		if dp[i] != -1 {
			return dp[i]
		}

		// Current index itself counts as 1
		ans := 1

		// Move right
		for j := i + 1; j <= min(n-1, i+d); j++ {

			// Stop if blocked
			if arr[j] >= arr[i] {
				break
			}

			// Update best answer
			ans = max(ans, 1+dfs(j))
		}

		// Move left
		for j := i - 1; j >= max(0, i-d); j-- {

			// Stop if blocked
			if arr[j] >= arr[i] {
				break
			}

			// Update best answer
			ans = max(ans, 1+dfs(j))
		}

		// Store result
		dp[i] = ans

		return ans
	}

	answer := 1

	// Try every starting index
	for i := 0; i < n; i++ {
		answer = max(answer, dfs(i))
	}

	return answer
}

// Helper function for maximum
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function for minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}