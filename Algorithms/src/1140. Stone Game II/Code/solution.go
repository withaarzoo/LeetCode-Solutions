func stoneGameII(piles []int) int {
	n := len(piles)

	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	var solve func(int, int) int

	solve = func(i, m int) int {
		if i == n {
			return 0
		}

		if dp[i][m] != -1 {
			return dp[i][m]
		}

		best := 0

		for x := 1; x <= 2*m && i+x <= n; x++ {
			nextM := m
			if x > nextM {
				nextM = x
			}

			current := suffix[i] - solve(i+x, nextM)
			if current > best {
				best = current
			}
		}

		dp[i][m] = best
		return best
	}

	return solve(0, 1)
}