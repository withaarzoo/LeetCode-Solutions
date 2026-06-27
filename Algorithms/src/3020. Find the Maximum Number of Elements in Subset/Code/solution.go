func maximumLength(nums []int) int {
	// Store frequency of every number
	freq := make(map[int64]int)

	for _, x := range nums {
		freq[int64(x)]++
	}

	ans := 1

	// Handle value 1 separately
	if cnt, ok := freq[1]; ok {
		// Only odd count of ones is valid
		if cnt%2 == 1 {
			if cnt > ans {
				ans = cnt
			}
		} else {
			if cnt-1 > ans {
				ans = cnt - 1
			}
		}
	}

	// Try every distinct starting value
	for start := range freq {
		if start == 1 {
			continue
		}

		cur := start
		length := 0

		for {
			cnt, ok := freq[cur]
			if !ok {
				break
			}

			// Use two copies if available
			if cnt >= 2 {
				length += 2

				// Move to the squared value
				cur = cur * cur
			} else {
				// Single copy becomes the center
				length++
				break
			}
		}

		// No center found
		if length%2 == 0 {
			length--
		}

		if length > ans {
			ans = length
		}
	}

	return ans
}