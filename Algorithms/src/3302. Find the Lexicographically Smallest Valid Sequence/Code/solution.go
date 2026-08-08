func validSequence(word1 string, word2 string) []int {
	n := len(word1) // I store the length of word1 for the two scans.
	m := len(word2) // I store the length of word2 because the answer has m indices.

	last := make([]int, m) // I store a position that can match every suffix of word2.

	// I initialize every position to -1 so I can detect missing matches.
	for i := 0; i < m; i++ {
		last[i] = -1 // I use -1 to mean that no matching position was found.
	}

	i := n - 1 // I start from the end of word1.
	j := m - 1 // I start matching from the end of word2.

	// I greedily match word2 from right to left.
	// This tells me where the remaining suffix can be matched.
	for i >= 0 && j >= 0 {
		// If the characters match, this index can represent word2[j].
		if word1[i] == word2[j] {
			last[j] = i // I remember this position for the suffix.
			j--         // I now need to match the previous character.
		}

		i-- // I continue moving toward the beginning of word1.
	}

	ans := make([]int, 0, m) // I reserve space for all m answer indices.
	canSkip := true           // I have not used the one allowed mismatch yet.
	j = 0                     // I start matching word2 from its first character.

	// I scan from left to right so that I always choose the earliest valid index.
	for i = 0; i < n && j < m; i++ {
		// If the current characters match, I can safely choose this index.
		if word1[i] == word2[j] {
			ans = append(ans, i) // I add the earliest matching index.
			j++                  // I move to the next character of word2.
		} else if canSkip && (j == m-1 || i < last[j+1]) {
			// I use the mismatch only when the remaining suffix can still be matched.
			canSkip = false      // I spend the only allowed mismatch.
			ans = append(ans, i) // I choose this earliest possible index.
			j++                  // I move to the next character of word2.
		}
	}

	// I return the answer only when every character of word2 was matched.
	if j == m {
		return ans
	}

	// If some character is still unmatched, no valid sequence exists.
	return []int{}
}