package main

import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	// sort by end asc, start desc if ends equal
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	const INF = -1 << 60
	a, b := INF, INF
	ans := 0

	for _, iv := range intervals {
		l, r := iv[0], iv[1]
		if l > b {
			ans += 2
			a = r - 1
			b = r
		} else if l > a {
			ans += 1
			a = b
			b = r
		} else {
			// both already inside
		}
	}
	return ans
}
