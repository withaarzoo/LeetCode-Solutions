package main

import (
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)

	// frequency map
	freq := make(map[int]int, n)
	for _, x := range nums {
		freq[x]++
	}

	ans := 1

	// Case A: existing values as target
	for v, already := range freq {
		lowVal := v - k
		highVal := v + k
		L := lowerBound(nums, lowVal)
		R := upperBound(nums, highVal)
		totalInRange := R - L
		need := totalInRange - already
		canFix := need
		if canFix > numOperations {
			canFix = numOperations
		}
		if already+canFix > ans {
			ans = already + canFix
		}
	}

	// Case B: sliding window for 2*k
	l := 0
	for r := 0; r < n; r++ {
		for l <= r && nums[r]-nums[l] > 2*k {
			l++
		}
		w := r - l + 1
		cand := w
		if cand > numOperations {
			cand = numOperations
		}
		if cand > ans {
			ans = cand
		}
	}

	return ans
}

func lowerBound(a []int, target int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func upperBound(a []int, target int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
