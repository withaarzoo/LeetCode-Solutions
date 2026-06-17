func processStr(s string, k int64) byte {
	n := len(s)

	// lengths[i] = length after processing s[i]
	lengths := make([]int64, n)
	var curLen int64 = 0

	for i := 0; i < n; i++ {
		c := s[i]

		if c >= 'a' && c <= 'z' {
			// Append a character
			curLen++
		} else if c == '*' {
			// Remove last character if present
			if curLen > 0 {
				curLen--
			}
		} else if c == '#' {
			// Duplicate the whole string
			curLen *= 2
		} else {
			// '%' only reverses, length unchanged
		}

		lengths[i] = curLen
	}

	// k is outside the final string
	if k >= curLen {
		return '.'
	}

	// Undo operations from right to left
	for i := n - 1; i >= 0; i-- {
		c := s[i]

		var before int64
		if i > 0 {
			before = lengths[i-1]
		}

		if c >= 'a' && c <= 'z' {
			// Letter was appended at index "before"
			if k == before {
				return c
			}
		} else if c == '#' {
			// Undo duplication
			if before > 0 {
				k %= before
			}
		} else if c == '%' {
			// Undo reverse
			k = before - 1 - k
		} else {
			// '*' keeps surviving indices unchanged
		}
	}

	return '.'
}