package main

func countPartitions(nums []int, k int) int {
	const MOD int64 = 1_000_000_007
	n := len(nums)

	dp := make([]int64, n+1)
	pref := make([]int64, n+1)

	dp[0] = 1
	pref[0] = 1

	maxdq := make([]int, 0) // indices
	mindq := make([]int, 0)

	l := 0

	for r := 0; r < n; r++ {
		x := nums[r]

		// maintain decreasing deque for max
		for len(maxdq) > 0 && nums[maxdq[len(maxdq)-1]] <= x {
			maxdq = maxdq[:len(maxdq)-1]
		}
		maxdq = append(maxdq, r)

		// maintain increasing deque for min
		for len(mindq) > 0 && nums[mindq[len(mindq)-1]] >= x {
			mindq = mindq[:len(mindq)-1]
		}
		mindq = append(mindq, r)

		// shrink window until valid
		for len(maxdq) > 0 && len(mindq) > 0 &&
			int64(nums[maxdq[0]]-nums[mindq[0]]) > int64(k) {
			if maxdq[0] == l {
				maxdq = maxdq[1:]
			}
			if mindq[0] == l {
				mindq = mindq[1:]
			}
			l++
		}

		L := l
		i := r + 1

		ways := pref[i-1]
		if L > 0 {
			ways -= pref[L-1]
		}
		ways %= MOD
		if ways < 0 {
			ways += MOD
		}

		dp[i] = ways
		pref[i] = (pref[i-1] + dp[i]) % MOD
	}

	return int(dp[n])
}
