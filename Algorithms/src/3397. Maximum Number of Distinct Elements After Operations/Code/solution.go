package main

import (
	"sort"
	"math"
)

func maxDistinctElements(nums []int, k int) int {
	n := len(nums)
	intervals := make([][2]int64, n)
	for i := 0; i < n; i++ {
		intervals[i][0] = int64(nums[i]) - int64(k)
		intervals[i][1] = int64(nums[i]) + int64(k)
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	lastAssigned := int64(math.MinInt64 / 4)
	ans := 0
	for _, iv := range intervals {
		l, r := iv[0], iv[1]
		assigned := l
		if lastAssigned+1 > assigned {
			assigned = lastAssigned + 1
		}
		if assigned <= r {
			ans++
			lastAssigned = assigned
		}
	}
	return ans
}
