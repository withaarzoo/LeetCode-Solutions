func maxValue(nums []int) []int {
	n := len(nums)

	// suffixMin[i] = smallest value in nums[i...n-1]
	// I keep one extra slot so the boundary check is simple.
	suffixMin := make([]int, n+1)
	suffixMin[n] = int(^uint(0) >> 1) // MaxInt for the current Go architecture

	for i := n - 1; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	ans := make([]int, n)
	l := 0

	// I process each connected component as one contiguous block.
	for l < n {
		r := l
		componentMax := nums[l]

		// I keep growing the block while an inversion crosses the next cut.
		for r+1 < n && componentMax > suffixMin[r+1] {
			r++
			if nums[r] > componentMax {
				componentMax = nums[r]
			}
		}

		// Every index in this block gets the same answer.
		for i := l; i <= r; i++ {
			ans[i] = componentMax
		}

		// Move to the next block.
		l = r + 1
	}

	return ans
}