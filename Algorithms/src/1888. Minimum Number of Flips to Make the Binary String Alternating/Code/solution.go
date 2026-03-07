func minFlips(s string) int {
	n := len(s)
	ss := s + s

	diff1 := 0
	diff2 := 0
	ans := 1<<31 - 1

	for i := 0; i < len(ss); i++ {

		expected1 := byte('0')
		expected2 := byte('1')

		if i%2 == 1 {
			expected1 = '1'
			expected2 = '0'
		}

		if ss[i] != expected1 {
			diff1++
		}
		if ss[i] != expected2 {
			diff2++
		}

		if i >= n {
			prev := ss[i-n]

			prevExp1 := byte('0')
			prevExp2 := byte('1')

			if (i-n)%2 == 1 {
				prevExp1 = '1'
				prevExp2 = '0'
			}

			if prev != prevExp1 {
				diff1--
			}
			if prev != prevExp2 {
				diff2--
			}
		}

		if i >= n-1 {
			if diff1 < ans {
				ans = diff1
			}
			if diff2 < ans {
				ans = diff2
			}
		}
	}

	return ans
}