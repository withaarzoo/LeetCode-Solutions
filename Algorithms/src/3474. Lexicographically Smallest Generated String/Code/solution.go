func generateString(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)
	length := n + m - 1

	ans := make([]byte, length)
	fixed := make([]bool, length)

	for i := 0; i < length; i++ {
		ans[i] = '?'
	}

	// Apply all 'T' constraints
	for i := 0; i < n; i++ {
		if str1[i] == 'T' {
			for j := 0; j < m; j++ {
				pos := i + j

				if ans[pos] != '?' && ans[pos] != str2[j] {
					return ""
				}

				ans[pos] = str2[j]
				fixed[pos] = true
			}
		}
	}

	// Fill remaining positions with 'a'
	for i := 0; i < length; i++ {
		if ans[i] == '?' {
			ans[i] = 'a'
		}
	}

	// Process all 'F' constraints
	for i := 0; i < n; i++ {
		if str1[i] == 'F' {
			same := true

			for j := 0; j < m; j++ {
				if ans[i+j] != str2[j] {
					same = false
					break
				}
			}

			if !same {
				continue
			}

			changed := false

			// Try changing from right to left
			for j := m - 1; j >= 0; j-- {
				pos := i + j

				if fixed[pos] {
					continue
				}

				for c := byte('a'); c <= byte('z'); c++ {
					if c != ans[pos] && c != str2[j] {
						ans[pos] = c
						changed = true
						break
					}
				}

				if changed {
					break
				}
			}

			if !changed {
				return ""
			}
		}
	}

	return string(ans)
}