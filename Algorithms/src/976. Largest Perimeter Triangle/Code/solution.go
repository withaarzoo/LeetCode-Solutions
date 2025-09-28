package main

import "sort"

// largestPerimeter returns the largest perimeter of a triangle
// that can be formed with three lengths from nums.
// Returns 0 if no non-degenerate triangle can be formed.
func largestPerimeter(nums []int) int {
    // Sort ascending
    sort.Ints(nums)
    n := len(nums)
    // Check triples from largest side downwards
    for i := n - 1; i >= 2; i-- {
        a := nums[i]     // largest in triple
        b := nums[i-1]
        c := nums[i-2]
        if b + c > a {
            return a + b + c
        }
    }
    return 0
}
