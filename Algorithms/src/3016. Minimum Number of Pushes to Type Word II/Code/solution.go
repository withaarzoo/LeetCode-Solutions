func minimumPushes(word string) int {
	// Store the frequency of every lowercase letter
	freq := make([]int, 26)

	// Count each character
	for _, ch := range word {
		freq[ch-'a']++
	}

	// Sort frequencies from largest to smallest
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	ans := 0

	// Assign push cost based on sorted position
	for i := 0; i < 26; i++ {
		// Stop once unused letters are reached
		if freq[i] == 0 {
			break
		}

		// Every 8 letters require one extra push
		pushes := (i / 8) + 1

		// Add this letter's total contribution
		ans += freq[i] * pushes
	}

	return ans
}