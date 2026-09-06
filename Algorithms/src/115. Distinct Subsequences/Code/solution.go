func numDistinct(s string, t string) int {
	m := len(t)

	// dp[j] stores the number of ways to form the first j characters of t.
	// There is exactly one way to form an empty string.
	dp := make([]uint64, m+1)
	dp[0] = 1

	// Process every character of s.
	for i := 0; i < len(s); i++ {
		c := s[i]

		// Traverse backwards so dp[j-1] still represents the previous state.
		// This prevents the current character from being reused.
		for j := m; j >= 1; j-- {
			// If the characters match, add the ways of forming the previous prefix.
			if c == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	// The problem guarantees that the answer fits in a 32-bit signed integer.
	return int(dp[m])
}