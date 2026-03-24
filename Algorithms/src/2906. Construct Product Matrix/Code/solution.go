func constructProductMatrix(grid [][]int) [][]int {
	const MOD int64 = 12345
	n, m := len(grid), len(grid[0])

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = 1
		}
	}

	var prefix int64 = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i][j] = int(prefix)
			prefix = (prefix * int64(grid[i][j])) % MOD
		}
	}

	var suffix int64 = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ans[i][j] = int((int64(ans[i][j]) * suffix) % MOD)
			suffix = (suffix * int64(grid[i][j])) % MOD
		}
	}

	return ans
}