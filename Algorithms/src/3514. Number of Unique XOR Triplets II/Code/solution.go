func uniqueXorTriplets(nums []int) int {
	const MAX_XOR = 2048

	// Store whether a value exists.
	present := make([]bool, MAX_XOR)
	for _, x := range nums {
		present[x] = true
	}

	// Initially only XOR = 0 is reachable.
	dp := make([]bool, MAX_XOR)
	dp[0] = true

	// Pick exactly 3 values.
	for step := 0; step < 3; step++ {
		next := make([]bool, MAX_XOR)

		// Extend every reachable XOR.
		for cur := 0; cur < MAX_XOR; cur++ {
			if !dp[cur] {
				continue
			}

			// Try every existing value.
			for v := 0; v < MAX_XOR; v++ {
				if present[v] {
					next[cur^v] = true
				}
			}
		}

		dp = next
	}

	// Count unique XOR values.
	ans := 0
	for _, ok := range dp {
		if ok {
			ans++
		}
	}

	return ans
}