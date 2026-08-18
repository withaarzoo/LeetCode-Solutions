func largestInteger(nums []int, k int) int {
	n := len(nums)

	// nums[i] is between 0 and 50, so a fixed-size frequency array is enough.
	// Its size is constant, so it does not grow with n.
	freq := [51]int{}

	// Count how many times every value occurs in nums.
	for _, x := range nums {
		freq[x]++
	}

	// When k = 1, every subarray contains exactly one element.
	// Therefore, only values occurring exactly once are valid.
	if k == 1 {
		// Check from 50 down to 0 to find the largest valid value first.
		for x := 50; x >= 0; x-- {
			if freq[x] == 1 {
				return x
			}
		}

		// No value occurs exactly once.
		return -1
	}

	// When k = n, the whole array is the only subarray.
	// Every distinct value therefore appears in exactly one subarray.
	if k == n {
		answer := 0

		// Find the largest value in nums.
		for _, x := range nums {
			if x > answer {
				answer = x
			}
		}

		return answer
	}

	// For 1 < k < n, only the first and last elements
	// can belong to exactly one subarray of size k.
	answer := -1

	// The first value is valid only if it occurs once in nums.
	if freq[nums[0]] == 1 && nums[0] > answer {
		answer = nums[0]
	}

	// The last value is valid only if it occurs once in nums.
	if freq[nums[n-1]] == 1 && nums[n-1] > answer {
		answer = nums[n-1]
	}

	// Return the largest valid endpoint, or -1 if none exists.
	return answer
}