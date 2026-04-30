func maxPathScore(grid [][]int, k int) int {
	m := len(grid)                         // Number of rows in the grid.
	n := len(grid[0])                      // Number of columns in the grid.
	const NEG = -1000000000               // Sentinel for impossible states.

	// prev[j][c] = best score at column j in the previous row using exact cost c.
	prev := make([][]int, n)
	for j := 0; j < n; j++ {
		prev[j] = make([]int, k+1)
		for c := 0; c <= k; c++ {
			prev[j][c] = NEG // Mark every state as impossible before the DP starts.
		}
	}

	for i := 0; i < m; i++ {
		// Rebuild the current row from scratch so old values do not leak into new states.
		curr := make([][]int, n)
		for j := 0; j < n; j++ {
			curr[j] = make([]int, k+1)
			for c := 0; c <= k; c++ {
				curr[j][c] = NEG // Reset every budget state for this row.
			}
		}

		for j := 0; j < n; j++ {
			gain := grid[i][j] // Score gained by taking this cell.
			need := 0
			if gain > 0 {
				need = 1 // 1 and 2 both spend one budget point.
			}

			limit := k
			if i+j < limit {
				limit = i + j // A path to (i, j) cannot spend more than i + j budget points.
			}

			// The starting cell is the base case.
			if i == 0 && j == 0 {
				curr[0][0] = 0 // Zero score and zero cost at the start.
				continue
			}

			for c := need; c <= limit; c++ {
				best := NEG

				// From above: use the completed previous row.
				if i > 0 && prev[j][c-need] != NEG {
					val := prev[j][c-need] + gain
					if val > best {
						best = val
					}
				}

				// From left: use the current row already computed.
				if j > 0 && curr[j-1][c-need] != NEG {
					val := curr[j-1][c-need] + gain
					if val > best {
						best = val
					}
				}

				curr[j][c] = best // Store the best exact-cost result for this cell.
			}
		}

		prev = curr // Move the current row into prev for the next iteration.
	}

	ans := NEG                        // Best score among all valid costs at the finish cell.
	for c := 0; c <= k; c++ {
		if prev[n-1][c] > ans {
			ans = prev[n-1][c]
		}
	}

	if ans < 0 {
		return -1 // If nothing is reachable, the answer is invalid.
	}
	return ans
}