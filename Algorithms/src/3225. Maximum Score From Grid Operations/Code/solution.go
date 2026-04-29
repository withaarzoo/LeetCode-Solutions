func maximumScore(grid [][]int) int64 {
	n := len(grid)
	if n == 1 {
		return 0
	}

	// pref[c][k] = sum of first k cells in column c
	pref := make([][]int64, n)
	for c := 0; c < n; c++ {
		pref[c] = make([]int64, n+1)
		var s int64 = 0
		for r := 0; r < n; r++ {
			s += int64(grid[r][c])
			pref[c][r+1] = s
		}
	}

	const NEG int64 = -(1 << 60)

	// dp[a][b] = best score after processing up to current column,
	// with previous height = a and current height = b.
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			val := pref[0][j] - pref[0][i]
			if val < 0 {
				val = 0
			}
			dp[i][j] = val
		}
	}

	for col := 1; col < n; col++ {
		ndp := make([][]int64, n+1)
		for i := 0; i <= n; i++ {
			ndp[i] = make([]int64, n+1)
			for j := 0; j <= n; j++ {
				ndp[i][j] = NEG
			}
		}

		for mid := 0; mid <= n; mid++ {
			q := make([]int64, n+1)
			for x := 0; x <= n; x++ {
				val := pref[col][x] - pref[col][mid]
				if val < 0 {
					val = 0
				}
				q[x] = val
			}

			prefixBest := make([]int64, n+1)
			prefixBest[0] = dp[0][mid]
			for a := 1; a <= n; a++ {
				if dp[a][mid] > prefixBest[a-1] {
					prefixBest[a] = dp[a][mid]
				} else {
					prefixBest[a] = prefixBest[a-1]
				}
			}

			suffixBest := make([]int64, n+2)
			for i := 0; i <= n+1; i++ {
				suffixBest[i] = NEG
			}
			suffixBest[n] = dp[n][mid] + q[n]
			for a := n - 1; a >= 0; a-- {
				cand := dp[a][mid] + q[a]
				if cand > suffixBest[a+1] {
					suffixBest[a] = cand
				} else {
					suffixBest[a] = suffixBest[a+1]
				}
			}

			limit := n
			if col == n-1 {
				limit = 0
			}

			for nxt := 0; nxt <= limit; nxt++ {
				best := NEG

				if prefixBest[nxt] != NEG {
					cand := prefixBest[nxt] + q[nxt]
					if cand > best {
						best = cand
					}
				}
				if suffixBest[nxt+1] != NEG && suffixBest[nxt+1] > best {
					best = suffixBest[nxt+1]
				}

				if best > ndp[mid][nxt] {
					ndp[mid][nxt] = best
				}
			}
		}

		dp = ndp
	}

	var ans int64 = 0
	for a := 0; a <= n; a++ {
		for b := 0; b <= n; b++ {
			if dp[a][b] > ans {
				ans = dp[a][b]
			}
		}
	}
	return ans
}