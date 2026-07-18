// Function to find GCD using the Euclidean Algorithm
func gcd(a, b int) int {
	// Continue until b becomes 0
	for b != 0 {
		a, b = b, a%b
	}

	// a now stores the GCD
	return a
}

func findGCD(nums []int) int {

	// Initialize minimum and maximum with the first element
	minimum := nums[0]
	maximum := nums[0]

	// Find the smallest and largest values
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}

	// Return the GCD of the smallest and largest numbers
	return gcd(minimum, maximum)
}