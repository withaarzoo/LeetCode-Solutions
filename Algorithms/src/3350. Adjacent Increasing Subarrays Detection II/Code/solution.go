func maxIncreasingSubarrays(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }

    inc := make([]int, n)
    inc[n-1] = 1
    for i := n - 2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            inc[i] = inc[i+1] + 1
        } else {
            inc[i] = 1
        }
    }

    feasible := func(k int) bool {
        if k == 0 {
            return true
        }
        for a := 0; a+2*k <= n; a++ {
            if inc[a] >= k && inc[a+k] >= k {
                return true
            }
        }
        return false
    }

    lo, hi, ans := 0, n/2, 0
    for lo <= hi {
        mid := (lo + hi) / 2
        if feasible(mid) {
            ans = mid
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return ans
}
