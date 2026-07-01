func numberOfSubstrings(s string) int {

	// Store the frequency of 'a', 'b', and 'c'
	freq := make([]int, 3)

	left := 0
	ans := 0
	n := len(s)

	// Expand the window
	for right := 0; right < n; right++ {

		// Add the current character
		freq[s[right]-'a']++

		// Shrink while all three characters exist
		for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {

			// Every larger ending position is also valid
			ans += n - right

			// Remove the leftmost character
			freq[s[left]-'a']--

			// Move the left pointer
			left++
		}
	}

	// Return the final answer
	return ans
}