func xorAfterQueries(nums []int, queries [][]int) int {
	const MOD int64 = 1_000_000_007

	n := len(nums)
	limit := int(math.Sqrt(float64(n))) + 1

	// Fast exponentiation
	modPow := func(base, exp int64) int64 {
		result := int64(1)

		for exp > 0 {
			if exp&1 == 1 {
				result = (result * base) % MOD
			}

			base = (base * base) % MOD
			exp >>= 1
		}

		return result
	}

	// Modular inverse
	var modInverse func(int64) int64
	modInverse = func(x int64) int64 {
		return modPow(x, MOD-2)
	}

	smallQueries := make(map[int][][]int)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]

		// Large k -> process directly
		if k >= limit {
			for i := l; i <= r; i += k {
				nums[i] = int((int64(nums[i]) * int64(v)) % MOD)
			}
		} else {
			smallQueries[k] = append(smallQueries[k], q)
		}
	}

	for k, group := range smallQueries {
		diff := make([]int64, n)

		for i := 0; i < n; i++ {
			diff[i] = 1
		}

		for _, q := range group {
			l, r, v := q[0], q[1], q[3]

			diff[l] = (diff[l] * int64(v)) % MOD

			steps := (r - l) / k
			nextPos := l + (steps+1)*k

			if nextPos < n {
				diff[nextPos] = (diff[nextPos] * modInverse(int64(v))) % MOD
			}
		}

		for i := 0; i < n; i++ {
			if i >= k {
				diff[i] = (diff[i] * diff[i-k]) % MOD
			}

			nums[i] = int((int64(nums[i]) * diff[i]) % MOD)
		}
	}

	answer := 0

	for _, num := range nums {
		answer ^= num
	}

	return answer
}