func subsequencePairCount(nums []int) int {
	const MOD = 1000000007
	const MAX = 200

	// Computes GCD using Euclid's algorithm.
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Current DP table.
	dp := make([][]int, MAX+1)
	for i := range dp {
		dp[i] = make([]int, MAX+1)
	}
	dp[0][0] = 1

	for _, x := range nums {
		// Next DP table after processing current number.
		next := make([][]int, MAX+1)
		for i := range next {
			next[i] = make([]int, MAX+1)
		}

		for g1 := 0; g1 <= MAX; g1++ {
			for g2 := 0; g2 <= MAX; g2++ {
				if dp[g1][g2] == 0 {
					continue
				}

				ways := dp[g1][g2]

				// Choice 1: Skip current number.
				next[g1][g2] = (next[g1][g2] + ways) % MOD

				// Choice 2: Put into first subsequence.
				ng1 := x
				if g1 != 0 {
					ng1 = gcd(g1, x)
				}
				next[ng1][g2] = (next[ng1][g2] + ways) % MOD

				// Choice 3: Put into second subsequence.
				ng2 := x
				if g2 != 0 {
					ng2 = gcd(g2, x)
				}
				next[g1][ng2] = (next[g1][ng2] + ways) % MOD
			}
		}

		dp = next
	}

	ans := 0

	// Count states where both GCDs are equal and non-zero.
	for g := 1; g <= MAX; g++ {
		ans = (ans + dp[g][g]) % MOD
	}

	return ans
}