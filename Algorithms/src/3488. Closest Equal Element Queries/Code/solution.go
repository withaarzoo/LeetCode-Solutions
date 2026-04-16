func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	// Store all indices for every value
	positions := make(map[int][]int)

	for i, num := range nums {
		positions[num] = append(positions[num], i)
	}

	// answer[i] = minimum circular distance for index i
	answer := make([]int, n)

	for i := 0; i < n; i++ {
		answer[i] = -1
	}

	// Process each value group
	for _, pos := range positions {
		m := len(pos)

		if m == 1 {
			continue
		}

		for i := 0; i < m; i++ {
			curr := pos[i]

			prev := pos[(i-1+m)%m]
			next := pos[(i+1)%m]

			distPrev := abs(curr - prev)
			distPrev = min(distPrev, n-distPrev)

			distNext := abs(curr - next)
			distNext = min(distNext, n-distNext)

			answer[curr] = min(distPrev, distNext)
		}
	}

	// Build final result
	result := make([]int, len(queries))

	for i, idx := range queries {
		result[i] = answer[idx]
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}