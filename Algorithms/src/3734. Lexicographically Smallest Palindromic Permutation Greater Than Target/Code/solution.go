func lexPalindromicPermutation(s string, target string) string {
	// Count the frequency of every lowercase English letter.
	frequency := make([]int, 26)
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}

	// Find the optional middle character and count odd frequencies.
	middle := byte(0)
	oddCount := 0
	for c := 0; c < 26; c++ {
		if frequency[c]%2 == 1 {
			oddCount++
			middle = byte('a' + c)
		}
	}

	// A palindrome can have at most one odd-frequency character.
	if oddCount > 1 {
		return ""
	}

	// Build the frequency multiset for the first half.
	halfCount := make([]int, 26)
	for c := 0; c < 26; c++ {
		halfCount[c] = frequency[c] / 2
	}

	k := len(s) / 2
	targetHalf := target[:k]

	// Find the smallest first-half permutation greater than or equal to targetHalf.
	smallestGreaterOrEqual := func(originalCount []int, prefix string) string {
		// Copy counts because this search modifies them.
		count := make([]int, 26)
		copy(count, originalCount)

		matched := 0

		// Match the target prefix exactly while the required character exists.
		for matched < k && count[prefix[matched]-'a'] > 0 {
			count[prefix[matched]-'a']--
			matched++
		}

		// The exact target prefix can be formed.
		if matched == k {
			return prefix
		}

		// Move backward until one position can be increased.
		for pos := matched; pos >= 0; pos-- {
			// Restore a character when backtracking over a matched position.
			if pos < matched {
				count[prefix[pos]-'a']++
			}

			// Choose the smallest available character larger than prefix[pos].
			for c := int(prefix[pos]-'a') + 1; c < 26; c++ {
				if count[c] == 0 {
					continue
				}

				// Keep the earlier prefix unchanged.
				result := make([]byte, 0, k)
				result = append(result, prefix[:pos]...)
				result = append(result, byte('a'+c))
				count[c]--

				// Fill the suffix in ascending order to minimize the result.
				for ch := 0; ch < 26; ch++ {
					for times := 0; times < count[ch]; times++ {
						result = append(result, byte('a'+ch))
					}
				}

				return string(result)
			}
		}

		return ""
	}

	// Build the complete palindrome by mirroring the first half.
	buildPalindrome := func(half string) string {
		result := make([]byte, 0, len(s))
		result = append(result, half...)

		if middle != 0 {
			result = append(result, middle)
		}

		// Append the left half in reverse order.
		for i := len(half) - 1; i >= 0; i-- {
			result = append(result, half[i])
		}

		return string(result)
	}

	// Change a byte slice into its next lexicographical permutation.
	nextPermutation := func(chars []byte) bool {
		pivot := len(chars) - 2

		// Find the rightmost position that can be increased.
		for pivot >= 0 && chars[pivot] >= chars[pivot+1] {
			pivot--
		}

		// No larger permutation exists.
		if pivot < 0 {
			return false
		}

		swapPos := len(chars) - 1

		// Find the smallest larger character in the suffix.
		for chars[swapPos] <= chars[pivot] {
			swapPos--
		}

		// Swap to make the smallest possible increase.
		chars[pivot], chars[swapPos] = chars[swapPos], chars[pivot]

		// Reverse the suffix to make it as small as possible.
		left, right := pivot+1, len(chars)-1
		for left < right {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}

		return true
	}

	// Find the smallest possible first half that can reach targetHalf.
	half := smallestGreaterOrEqual(halfCount, targetHalf)

	if half == "" && k > 0 {
		return ""
	}

	// Build the smallest candidate from this first half.
	candidate := buildPalindrome(half)

	// If it is already strictly greater, it is the answer.
	if candidate > target {
		return candidate
	}

	// Otherwise, move to the next larger first-half permutation.
	chars := []byte(half)
	if !nextPermutation(chars) {
		return ""
	}

	// The next first half gives the lexicographically smallest valid answer.
	return buildPalindrome(string(chars))
}