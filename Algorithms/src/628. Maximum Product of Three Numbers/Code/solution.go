func maximumProduct(nums []int) int {

	// Store the three largest numbers
	max1, max2, max3 := -1001, -1001, -1001

	// Store the two smallest numbers
	min1, min2 := 1001, 1001

	// Traverse the array once
	for _, num := range nums {

		// Update the three largest numbers
		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num >= max3 {
			max3 = num
		}

		// Update the two smallest numbers
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}
	}

	// Product of the three largest numbers
	product1 := max1 * max2 * max3

	// Product of the two smallest numbers and the largest number
	product2 := min1 * min2 * max1

	// Return the larger product
	if product1 > product2 {
		return product1
	}
	return product2
}