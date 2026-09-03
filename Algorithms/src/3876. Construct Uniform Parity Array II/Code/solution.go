func uniformArray(nums1 []int) bool {
	// Use a value larger than every possible input as the initial minimum.
	const inf = int(1e18)

	// Store the smallest odd and even values found in the array.
	minOdd := inf
	minEven := inf

	// Scan the array once because only the two minimum values are needed.
	for _, x := range nums1 {
		if x%2 == 0 {
			// Keep the smallest even value because it is the hardest to change.
			if x < minEven {
				minEven = x
			}
		} else {
			// Keep the smallest odd value because it is the best subtraction value.
			if x < minOdd {
				minOdd = x
			}
		}
	}

	// If no odd value exists, all numbers are already even.
	if minOdd == inf {
		return true
	}

	// Every even value must be larger than the smallest odd value.
	// Then subtracting minOdd changes each even value into an odd value.
	return minOdd < minEven
}