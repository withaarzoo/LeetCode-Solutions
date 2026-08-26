func shortestBeautifulSubstring(s string, k int) string {
	answer := "" // I store the best beautiful substring found so far.
	left := 0    // I keep the left boundary of the sliding window.
	ones := 0    // I count how many '1' characters are inside the window.

	// I expand the window by moving the right pointer through the string.
	for right := 0; right < len(s); right++ {
		// I update the count when the newly added character is '1'.
		if s[right] == '1' {
			ones++
		}

		// If I have too many ones, I remove characters from the left
		// until the window contains at most k ones again.
		for ones > k {
			if s[left] == '1' {
				ones--
			}
			left++
		}

		// I remove unnecessary leading zeros because they do not change
		// the number of ones and only make the valid substring longer.
		for ones == k && s[left] == '0' {
			left++
		}

		// The current window is beautiful when it contains exactly k ones.
		if ones == k {
			// I extract the shortest valid candidate ending at right.
			candidate := s[left : right+1]

			// I update the answer if this candidate is shorter, or if equal
			// lengths require lexicographical comparison.
			if answer == "" ||
				len(candidate) < len(answer) ||
				(len(candidate) == len(answer) && candidate < answer) {
				answer = candidate
			}
		}
	}

	// I return the best substring, or an empty string if none exists.
	return answer
}