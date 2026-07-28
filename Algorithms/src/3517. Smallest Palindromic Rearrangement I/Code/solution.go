func smallestPalindrome(s string) string {

	// Store frequency of every lowercase letter
	freq := make([]int, 26)

	// Count character frequencies
	for _, ch := range s {
		freq[ch-'a']++
	}

	left := make([]byte, 0)
	middle := byte(0)

	// Build the left half and find the middle character
	for i := 0; i < 26; i++ {

		// Add half of the occurrences to the left side
		for j := 0; j < freq[i]/2; j++ {
			left = append(left, byte('a'+i))
		}

		// Odd frequency character becomes the center
		if freq[i]%2 == 1 {
			middle = byte('a' + i)
		}
	}

	// Right half is the reverse of the left half
	right := make([]byte, len(left))
	for i := 0; i < len(left); i++ {
		right[i] = left[len(left)-1-i]
	}

	// Build the final answer
	if middle != 0 {
		return string(left) + string(middle) + string(right)
	}

	return string(left) + string(right)
}