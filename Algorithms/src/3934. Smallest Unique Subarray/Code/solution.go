func smallestUniqueSubarray(nums []int) int {
    n := len(nums)
    const MOD1, MOD2 = 1000000007, 1000000009
    const BASE1, BASE2 = 1000003, 1000033
    h1 := make([]int64, n+1)
    h2 := make([]int64, n+1)
    p1 := make([]int64, n+1)
    p2 := make([]int64, n+1)
    p1[0], p2[0] = 1, 1
    for i := 0; i < n; i++ {
        h1[i+1] = (h1[i]*BASE1 + int64(nums[i])) % MOD1
        h2[i+1] = (h2[i]*BASE2 + int64(nums[i])) % MOD2
        p1[i+1] = p1[i] * BASE1 % MOD1
        p2[i+1] = p2[i] * BASE2 % MOD2
    }
    hasUnique := func(length int) bool {
        cnt := make(map[[2]int64]int)
        for i := 0; i+length <= n; i++ {
            a := (h1[i+length] - h1[i]*p1[length]%MOD1 + MOD1) % MOD1
            b := (h2[i+length] - h2[i]*p2[length]%MOD2 + MOD2) % MOD2
            cnt[[2]int64{a, b}]++
        }
        for _, v := range cnt {
            if v == 1 {
                return true
            }
        }
        return false
    }
    lo, hi, ans := 1, n, n
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if hasUnique(mid) {
            ans = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return ans
}