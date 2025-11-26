package main

func numberOfPaths(grid [][]int, k int) int {
	const MOD int = 1_000_000_007
	m, n := len(grid), len(grid[0])

	// prev[j][r] = paths to (i-1, j) with sum % k == r
	// cur[j][r]  = paths to (i,   j) with sum % k == r
	prev := make([][]int, n)
	cur := make([][]int, n)
	for j := 0; j < n; j++ {
		prev[j] = make([]int, k)
		cur[j] = make([]int, k)
	}

	for i := 0; i < m; i++ {
		// reset current row
		for j := 0; j < n; j++ {
			for r := 0; r < k; r++ {
				cur[j][r] = 0
			}
		}

		for j := 0; j < n; j++ {
			val := grid[i][j] % k

			// starting cell
			if i == 0 && j == 0 {
				cur[0][val] = 1
				continue
			}

			// from top
			if i > 0 {
				for r := 0; r < k; r++ {
					if prev[j][r] == 0 {
						continue
					}
					nr := (r + val) % k
					cur[j][nr] += prev[j][r]
					if cur[j][nr] >= MOD {
						cur[j][nr] -= MOD
					}
				}
			}

			// from left
			if j > 0 {
				for r := 0; r < k; r++ {
					if cur[j-1][r] == 0 {
						continue
					}
					nr := (r + val) % k
					cur[j][nr] += cur[j-1][r]
					if cur[j][nr] >= MOD {
						cur[j][nr] -= MOD
					}
				}
			}
		}

		// swap prev and cur
		prev, cur = cur, prev
	}

	return prev[n-1][0]
}
