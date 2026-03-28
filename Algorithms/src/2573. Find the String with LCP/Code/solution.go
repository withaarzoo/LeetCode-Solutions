func findTheString(lcp [][]int) string {
	n := len(lcp)

	// group[i] = which character group position i belongs to
	group := make([]int, n)
	for i := 0; i < n; i++ {
		group[i] = -1
	}

	curGroup := 0

	// Build groups
	for i := 0; i < n; i++ {
		if group[i] == -1 {
			if curGroup == 26 {
				return ""
			}

			group[i] = curGroup
			curGroup++

			for j := i + 1; j < n; j++ {
				if lcp[i][j] > 0 {
					group[j] = group[i]
				}
			}
		}
	}

	// Build answer string
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = byte('a' + group[i])
	}

	// Verify using DP
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if ans[i] == ans[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			}
		}
	}

	// Compare matrices
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(ans)
}