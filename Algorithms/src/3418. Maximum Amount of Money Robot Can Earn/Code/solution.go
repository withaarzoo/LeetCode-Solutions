func maximumAmount(coins [][]int) int {
	m := len(coins)
	n := len(coins[0])
	NEG := -1000000000

	// dp[i][j][k] = max money at (i,j) using k neutralizations
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				dp[i][j][k] = NEG
			}
		}
	}

	// Starting cell
	if coins[0][0] >= 0 {
		dp[0][0][0] = coins[0][0]
	} else {
		dp[0][0][0] = coins[0][0]
		dp[0][0][1] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= 2; k++ {
				if dp[i][j][k] == NEG {
					continue
				}

				// Move Down
				if i+1 < m {
					val := coins[i+1][j]

					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k]+val)

					if val < 0 && k < 2 {
						dp[i+1][j][k+1] = max(dp[i+1][j][k+1], dp[i][j][k])
					}
				}

				// Move Right
				if j+1 < n {
					val := coins[i][j+1]

					dp[i][j+1][k] = max(dp[i][j+1][k], dp[i][j][k]+val)

					if val < 0 && k < 2 {
						dp[i][j+1][k+1] = max(dp[i][j+1][k+1], dp[i][j][k])
					}
				}
			}
		}
	}

	ans := dp[m-1][n-1][0]
	ans = max(ans, dp[m-1][n-1][1])
	ans = max(ans, dp[m-1][n-1][2])

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}