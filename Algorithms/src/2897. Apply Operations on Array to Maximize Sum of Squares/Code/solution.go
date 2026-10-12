func maxSum(nums []int, k int) int {
    const MOD = 1000000007
    n := len(nums)
    cnt := make([]int, 30)
    for _, x := range nums {
        for b := 0; b < 30; b++ {
            if x&(1<<b) != 0 {
                cnt[b]++
            }
        }
    }
    a := make([]int64, n)
    for b := 29; b >= 0; b-- {
        c := cnt[b]
        for i := 0; i < c; i++ {
            a[i] |= 1 << b
        }
    }
    var ans int64
    for i := 0; i < k; i++ {
        ans = (ans + a[i]*a[i]) % MOD
    }
    return int(ans)
}