func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	ans := 0

	// Try every possible starting index
	for left := 0; left < n; left++ {
		countTarget := 0

		// Extend the subarray
		for right := left; right < n; right++ {

			// Update target frequency
			if nums[right] == target {
				countTarget++
			}

			// Current subarray length
			length := right - left + 1

			// Check majority condition
			if 2*countTarget > length {
				ans++
			}
		}
	}

	return ans
}