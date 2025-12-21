func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	sorted := make([]bool, n-1)
	deletions := 0

	for col := 0; col < m; col++ {
		needDelete := false

		for row := 0; row < n-1; row++ {
			if !sorted[row] && strs[row][col] > strs[row+1][col] {
				needDelete = true
				break
			}
		}

		if needDelete {
			deletions++
			continue
		}

		for row := 0; row < n-1; row++ {
			if !sorted[row] && strs[row][col] < strs[row+1][col] {
				sorted[row] = true
			}
		}
	}

	return deletions
}
