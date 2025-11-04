package main

import (
	"sort"
)

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	ans := make([]int, 0, n-k+1)
	freq := map[int]int{}

	// initial window
	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}

	ans = append(ans, computeXSum(freq, x))

	// slide
	for i := k; i < n; i++ {
		add := nums[i]
		rem := nums[i-k]

		freq[add]++
		freq[rem]--
		if freq[rem] == 0 {
			delete(freq, rem)
		}
		ans = append(ans, computeXSum(freq, x))
	}

	return ans
}

type item struct {
	val int
	fr  int
}

func computeXSum(freq map[int]int, x int) int {
	items := make([]item, 0, len(freq))
	for v, f := range freq {
		items = append(items, item{v, f})
	}
	// sort by freq desc, value desc
	sort.Slice(items, func(i, j int) bool {
		if items[i].fr != items[j].fr {
			return items[i].fr > items[j].fr
		}
		return items[i].val > items[j].val
	})
	sum := 0
	take := x
	if take > len(items) {
		take = len(items)
	}
	for i := 0; i < take; i++ {
		sum += items[i].val * items[i].fr
	}
	return sum
}
