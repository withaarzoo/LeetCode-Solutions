func isTrionic(nums []int) bool {
	n := len(nums)
	i := 0

	// 1) strictly increasing
	for i+1 < n && nums[i] < nums[i+1] {
		i++
	}
	if i == 0 || i == n-1 {
		return false
	}

	// 2) strictly decreasing
	mid := i
	for i+1 < n && nums[i] > nums[i+1] {
		i++
	}
	if i == mid || i == n-1 {
		return false
	}

	// 3) strictly increasing again
	for i+1 < n && nums[i] < nums[i+1] {
		i++
	}

	return i == n-1
}
