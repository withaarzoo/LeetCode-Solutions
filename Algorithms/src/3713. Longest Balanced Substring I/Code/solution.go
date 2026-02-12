func longestBalanced(s string) int {
	n := len(s)
	ans := 0

	for i := 0; i < n; i++ {
		freq := make([]int, 26)
		distinct := 0
		maxFreq := 0

		for j := i; j < n; j++ {
			idx := int(s[j] - 'a')

			if freq[idx] == 0 {
				distinct++
			}

			freq[idx]++
			if freq[idx] > maxFreq {
				maxFreq = freq[idx]
			}

			length := j - i + 1

			if length == distinct*maxFreq {
				if length > ans {
					ans = length
				}
			}
		}
	}

	return ans
}
