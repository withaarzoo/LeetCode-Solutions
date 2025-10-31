package main

func getSneakyNumbers(nums []int) []int {
    n := len(nums) - 2           // original range size
    seen := make([]bool, n)      // marks for 0..n-1
    res := make([]int, 0, 2)
    for _, x := range nums {
        if seen[x] {
            res = append(res, x) // this value is repeated
            if len(res) == 2 {
                break            // both found, stop early
            }
        } else {
            seen[x] = true      // mark seen first time
        }
    }
    return res
}
