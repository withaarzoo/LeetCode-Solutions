func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int) // Stores the frequency of each value inside the current window.
	left := 0                  // Left boundary of the sliding window.
	ans := 0                   // Stores the longest valid window length.

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++ // Add nums[right] to the window and increase its frequency.

		for freq[nums[right]] > k {
			freq[nums[left]]-- // Remove nums[left] because the current window is invalid.
			left++             // Move the left boundary forward.
		}

		if right-left+1 > ans {
			ans = right - left + 1 // Update the maximum length using the valid window.
		}
	}

	return ans // Return the length of the longest valid subarray.
}