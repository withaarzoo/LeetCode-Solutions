func uniqueXorTriplets(nums []int) int {
	// Length of the permutation
	n := len(nums)

	// Handle the small cases
	if n <= 2 {
		return n
	}

	// Count how many bits are needed to represent n
	bits := 0
	x := n
	for x > 0 {
		bits++
		x >>= 1
	}

	// Every value in [0, 2^bits - 1] can be generated
	return 1 << bits
}