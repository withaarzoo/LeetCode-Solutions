package main

// countValidSelections simulates the process for each zero index and both directions
func simulate(nums []int, start int, dir int) bool {
    n := len(nums)
    a := make([]int, n)
    copy(a, nums)
    curr := start
    for curr >= 0 && curr < n {
        if a[curr] == 0 {
            curr += dir // move same direction
        } else {
            a[curr]--    // decrement
            dir = -dir   // reverse direction
            curr += dir  // step in new direction
        }
    }
    for _, v := range a {
        if v != 0 {
            return false
        }
    }
    return true
}

func countValidSelections(nums []int) int {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            continue
        }
        if simulate(nums, i, -1) {
            ans++ // left
        }
        if simulate(nums, i, 1) {
            ans++ // right
        }
    }
    return ans
}
