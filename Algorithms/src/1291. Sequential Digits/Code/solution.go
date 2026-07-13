func sequentialDigits(low int, high int) []int {

	// String containing all consecutive digits
	digits := "123456789"

	// Result slice
	ans := []int{}

	// Number of digits in low and high
	minLen := len([]byte(fmt.Sprintf("%d", low)))
	maxLen := len([]byte(fmt.Sprintf("%d", high)))

	// Try every possible length
	for length := minLen; length <= maxLen; length++ {

		// Generate every substring of current length
		for start := 0; start+length <= 9; start++ {

			// Convert substring into an integer
			num, _ := strconv.Atoi(digits[start : start+length])

			// Keep only numbers inside the range
			if num >= low && num <= high {
				ans = append(ans, num)
			}
		}
	}

	return ans
}