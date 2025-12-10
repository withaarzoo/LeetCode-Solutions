package main

func countPermutations(complexity []int) int {
	const MOD int64 = 1_000_000_007
	n := len(complexity)

	// 1. Find global minimum and its count
	minVal := complexity[0]
	cntMin := 0
	for _, x := range complexity {
		if x < minVal {
			minVal = x
			cntMin = 1
		} else if x == minVal {
			cntMin++
		}
	}

	// 2. Check if index 0 has unique minimum
	if complexity[0] != minVal || cntMin != 1 {
		return 0
	}

	// 3. Compute (n - 1)! % MOD
	ans := int64(1)
	for i := 2; i <= n-1; i++ {
		ans = (ans * int64(i)) % MOD
	}
	return int(ans)
}
