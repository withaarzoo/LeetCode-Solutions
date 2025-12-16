func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	tree := make([][]int, n)
	for _, e := range hierarchy {
		tree[e[0]-1] = append(tree[e[0]-1], e[1]-1)
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, budget+1)
		}
	}

	merge := func(A, B []int) []int {
		C := make([]int, budget+1)
		for i := range C {
			C[i] = -1e9
		}
		for i := 0; i <= budget; i++ {
			if A[i] < 0 {
				continue
			}
			for j := 0; i+j <= budget; j++ {
				if C[i+j] < A[i]+B[j] {
					C[i+j] = A[i] + B[j]
				}
			}
		}
		return C
	}

	var dfs func(int)
	dfs = func(u int) {
		for _, v := range tree[u] {
			dfs(v)
		}

		for pb := 0; pb <= 1; pb++ {
			price := present[u]
			if pb == 1 {
				price /= 2
			}
			profit := future[u] - price

			skip := make([]int, budget+1)
			for _, v := range tree[u] {
				skip = merge(skip, dp[v][0])
			}

			take := make([]int, budget+1)
			for i := range take {
				take[i] = -1e9
			}

			if price <= budget {
				base := make([]int, budget+1)
				for _, v := range tree[u] {
					base = merge(base, dp[v][1])
				}
				for b := price; b <= budget; b++ {
					take[b] = base[b-price] + profit
				}
			}

			for b := 0; b <= budget; b++ {
				dp[u][pb][b] = max(skip[b], take[b])
			}
		}
	}

	dfs(0)
	ans := 0
	for _, v := range dp[0][0] {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
