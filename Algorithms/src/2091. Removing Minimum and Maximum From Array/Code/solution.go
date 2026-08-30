func minimumDeletions(nums []int) int {
	// Store the total number of elements in the array.
	n := len(nums)

	// Start by assuming the first element is both minimum and maximum.
	minIndex := 0
	maxIndex := 0

	// Find the positions of the minimum and maximum elements.
	for i := 1; i < n; i++ {
		// Update the minimum index if a smaller value is found.
		if nums[i] < nums[minIndex] {
			minIndex = i
		}

		// Update the maximum index if a larger value is found.
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	// Remove everything from the front up to the farther special element.
	removeFromFront := max(minIndex, maxIndex) + 1

	// Remove everything from the back up to the farther special element.
	removeFromBack := n - min(minIndex, maxIndex)

	// Calculate both ways of removing one element from each side.
	removeFromBothSides := min(
		minIndex+1+(n-maxIndex),
		maxIndex+1+(n-minIndex),
	)

	// Return the minimum deletions among all possible strategies.
	return min(
		removeFromFront,
		min(removeFromBack, removeFromBothSides),
	)
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}