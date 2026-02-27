func minOperations(s string, k int) int {
	n := len(s)
	zero := 0

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zero++
		}
	}

	if zero == 0 {
		return 0
	}

	if n == k {
		if zero == n {
			return 1
		}
		return -1
	}

	one := n - zero
	base := n - k

	ans := int(^uint(0) >> 1) // max int

	// Odd operations
	if (k%2) == (zero%2) {
		m := max(
			(zero+k-1)/k,
			(one+base-1)/base,
		)
		if m%2 == 0 {
			m++
		}
		if m < ans {
			ans = m
		}
	}

	// Even operations
	if zero%2 == 0 {
		m := max(
			(zero+k-1)/k,
			(zero+base-1)/base,
		)
		if m%2 == 1 {
			m++
		}
		if m < ans {
			ans = m
		}
	}

	if ans == int(^uint(0)>>1) {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}