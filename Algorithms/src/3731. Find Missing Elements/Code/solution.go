func findMissingElements(nums []int) []int {
	// Store every number for constant-time lookup
	seen := make(map[int]bool)

	// Initialize minimum and maximum
	mn, mx := nums[0], nums[0]

	// Fill the map and find the range
	for _, num := range nums {
		seen[num] = true

		if num < mn {
			mn = num
		}

		if num > mx {
			mx = num
		}
	}

	// Store missing numbers
	ans := []int{}

	// Check every value in the range
	for x := mn; x <= mx; x++ {
		// If the value is missing, add it
		if !seen[x] {
			ans = append(ans, x)
		}
	}

	return ans
}