func lexGreaterPermutation(s string, target string) string {
	// Store the frequency of every lowercase character available in s.
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	n := len(s)
	matched := 0

	// Match target from left to right for as long as possible.
	for matched < n && count[target[matched]-'a'] > 0 {
		// Using the same character keeps the current prefix smallest.
		count[target[matched]-'a']--
		matched++
	}

	// Start where matching failed, or at the last position if all matched.
	start := n - 1
	if matched < n {
		start = matched
	}

	// Try to increase the answer from right to left.
	for i := start; i >= 0; i-- {
		// Restore this character if it was previously used to match target.
		if i < matched {
			count[target[i]-'a']++
		}

		// Find the smallest available character strictly greater than target[i].
		bigger := -1
		for ch := int(target[i]-'a') + 1; ch < 26; ch++ {
			if count[ch] > 0 {
				bigger = ch
				break
			}
		}

		// Build the answer when an increasing character is found.
		if bigger != -1 {
			// Consume the character that makes the string strictly greater.
			count[bigger]--

			// Keep the prefix before i unchanged.
			answer := make([]byte, 0, n)
			answer = append(answer, target[:i]...)

			// Add the smallest available character greater than target[i].
			answer = append(answer, byte('a'+bigger))

			// Append all remaining characters in sorted order.
			for ch := 0; ch < 26; ch++ {
				for count[ch] > 0 {
					answer = append(answer, byte('a'+ch))
					count[ch]--
				}
			}

			return string(answer)
		}
	}

	// No permutation is strictly greater than target.
	return ""
}