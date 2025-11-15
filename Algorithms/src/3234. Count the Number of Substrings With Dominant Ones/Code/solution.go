package main

import (
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	var ans int64 = 0
	// all-ones substrings
	var run int64 = 0
	for i := 0; i < n; i++ {
		if s[i] == '1' { run++ }
		else { ans += run * (run + 1) / 2; run = 0 }
	}
	ans += run * (run + 1) / 2

	// zero positions
	zeroPos := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '0' { zeroPos = append(zeroPos, i) }
	}
	m := len(zeroPos)
	if m == 0 {
		return int(ans)
	}

	K := int(math.Floor(math.Sqrt(float64(n))))
	for k := 1; k <= K && k <= m; k++ {
		for i := 0; i + k - 1 < m; i++ {
			leftPrev := -1
			if i != 0 { leftPrev = zeroPos[i-1] }
			rightNext := n
			if i + k - 1 != m-1 { rightNext = zeroPos[i+k] }
			leftOnes := zeroPos[i] - leftPrev - 1
			rightOnes := rightNext - zeroPos[i+k-1] - 1
			baseLen := zeroPos[i+k-1] - zeroPos[i] + 1
			needLen := k*k + k
			t := needLen - baseLen
			totalPairs := int64((leftOnes + 1) * (rightOnes + 1))
			if t <= 0 { ans += totalPairs; continue }

			var pairs_lt int64 = 0
			s0 := int64(t - 1)
			if s0 >= 0 {
				L := int64(leftOnes)
				R := int64(rightOnes)
				x_max := L
				if s0 < x_max { x_max = s0 }
				if x_max >= 0 {
					x0 := int64(0)
					if s0 - R > 0 { x0 = s0 - R }
					if x0 > x_max {
						pairs_lt = (x_max + 1) * (R + 1)
					} else {
						part1 := x0 * (R + 1)
						n2 := x_max - x0 + 1
						sum_x := (x0 + x_max) * n2 / 2
						part2 := n2 * (s0 + 1) - sum_x
						pairs_lt = part1 + part2
					}
				} else {
					pairs_lt = 0
				}
			} else {
				pairs_lt = 0
			}

			valid := totalPairs - pairs_lt
			if valid > 0 { ans += valid }
		}
	}

	return int(ans)
}
