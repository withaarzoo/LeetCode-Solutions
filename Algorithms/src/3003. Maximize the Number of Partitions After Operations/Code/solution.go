package main

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	cache := make(map[[3]int]int)

	var bitCount = func(x int) int {
		count := 0
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return count
	}

	var dp func(int, int, int) int
	dp = func(i int, canChange int, mask int) int {
		if i == n {
			return 0
		}
		key := [3]int{i, canChange, mask}
		if val, ok := cache[key]; ok {
			return val
		}
		bit := int(s[i] - 'a')
		newMask := mask | (1 << bit)
		res := 0

		if bitCount(newMask) > k {
			res = 1 + dp(i+1, canChange, 1<<bit)
		} else {
			res = dp(i+1, canChange, newMask)
		}

		if canChange == 1 {
			for j := 0; j < 26; j++ {
				changed := mask | (1 << j)
				if bitCount(changed) > k {
					res = max(res, 1+dp(i+1, 0, 1<<j))
				} else {
					res = max(res, dp(i+1, 0, changed))
				}
			}
		}
		cache[key] = res
		return res
	}

	return dp(0, 1, 0) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
