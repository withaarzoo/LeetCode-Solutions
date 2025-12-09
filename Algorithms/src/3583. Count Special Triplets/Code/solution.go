package main

func specialTriplets(nums []int) int {
    const MOD int64 = 1_000_000_007

    // Maps to store frequencies on left and right
    right := make(map[int]int64)
    left := make(map[int]int64)

    for _, x := range nums {
        right[x]++
    }

    var ans int64 = 0

    for _, x := range nums {
        // x is now the middle, remove from right
        right[x]--

        target := x * 2

        cntLeft := left[target]
        cntRight := right[target]

        add := (cntLeft * cntRight) % MOD
        ans = (ans + add) % MOD

        // move x to left
        left[x]++
    }

    return int(ans % MOD)
}
