package main

import (
	"sort"
)

// countPairs counts the number of pairs (i, j) in the sorted array `nums`
// such that the difference between the pair (nums[j] - nums[i]) is less than or equal to `mid`.
func countPairs(nums []int, mid int) int {
	count := 0  // Initialize count to store the number of valid pairs
	j := 0      // Pointer to traverse the array and find valid pairs

	// Loop through each element in the array
	for i := 0; i < len(nums); i++ {
		// Move the pointer `j` to the right until the difference exceeds `mid`
		for j < len(nums) && nums[j]-nums[i] <= mid {
			j++
		}
		// Add the number of valid pairs formed with element at index `i`
		// Subtract 1 because we are only interested in pairs where j > i
		count += j - i - 1
	}
	return count  // Return the total count of valid pairs
}

// smallestDistancePair returns the k-th smallest distance among all pairs in the array `nums`.
func smallestDistancePair(nums []int, k int) int {
	// First, sort the array to make it easier to calculate distances between pairs
	sort.Ints(nums)

	// Initialize the binary search range for the smallest and largest possible distances
	low := 0                            // Smallest possible distance
	high := nums[len(nums)-1] - nums[0] // Largest possible distance

	// Perform binary search to find the k-th smallest distance
	for low < high {
		// Calculate the midpoint of the current range
		mid := (low + high) / 2
		
		// Count how many pairs have a distance less than or equal to `mid`
		if countPairs(nums, mid) >= k {
			// If there are at least `k` pairs with distance <= `mid`,
			// then the k-th smallest distance is <= `mid`
			high = mid
		} else {
			// Otherwise, the k-th smallest distance is greater than `mid`
			low = mid + 1
		}
	}
	// The binary search converges to the k-th smallest distance
	return low
}
