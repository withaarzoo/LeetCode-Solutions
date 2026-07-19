func smallestSubsequence(s string) string {

	// Count remaining occurrences of every character
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Track whether a character is already inside the stack
	inStack := make([]bool, 26)

	// Stack to build the answer
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {

		ch := s[i]

		// Current occurrence has been processed
		freq[ch-'a']--

		// Skip duplicate characters
		if inStack[ch-'a'] {
			continue
		}

		// Remove larger characters if they appear again later
		for len(stack) > 0 &&
			stack[len(stack)-1] > ch &&
			freq[stack[len(stack)-1]-'a'] > 0 {

			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		// Push current character
		stack = append(stack, ch)
		inStack[ch-'a'] = true
	}

	// Convert stack into the final answer
	return string(stack)
}