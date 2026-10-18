func numOfWays(nums []int) int {
    const MOD = 1000000007
    n := len(nums)
    fact := make([]int64, n+1)
    inv := make([]int64, n+1)
    fact[0] = 1
    for i := 1; i <= n; i++ {
        fact[i] = fact[i-1] * int64(i) % MOD
    }
    inv[n] = modPow(fact[n], MOD-2, MOD)
    for i := n - 1; i >= 0; i-- {
        inv[i] = inv[i+1] * int64(i+1) % MOD
    }
    var ways func([]int) int64
    ways = func(arr []int) int64 {
        m := len(arr)
        if m <= 2 {
            return 1
        }
        left, right := []int{}, []int{}
        for _, x := range arr {
            if x < arr[0] {
                left = append(left, x)
            } else if x > arr[0] {
                right = append(right, x)
            }
        }
        return comb(m-1, len(left), fact, inv, MOD) * ways(left) % MOD * ways(right) % MOD
    }
    return int((ways(nums) - 1 + MOD) % MOD)
}
func comb(a, b int, fact, inv []int64, MOD int64) int64 {
    if b < 0 || b > a {
        return 0
    }
    return fact[a] * inv[b] % MOD * inv[a-b] % MOD
}
func modPow(base, exp, mod int64) int64 {
    res := int64(1)
    for exp > 0 {
        if exp&1 == 1 {
            res = res * base % mod
        }
        base = base * base % mod
        exp >>= 1
    }
    return res
}